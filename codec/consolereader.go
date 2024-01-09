package codec

import (
	"bufio"
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	pbcodec "github.com/ChainSafe/firehose-arweave/pb/sf/arweave/type/v1"
	"github.com/dvsekhvalnov/jose2go/base64url"
	"github.com/streamingfast/bstream"
	pbbstream "github.com/streamingfast/bstream/pb/sf/bstream/v1"
	"github.com/streamingfast/dmetrics"
	firecore "github.com/streamingfast/firehose-core"
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

var supportedVersions = []uint64{13}
var supportedVersionStrings = []string{"13"}

// ConsoleReader is what reads the `geth` output directly. It builds
// up some LogEntry objects. See `LogReader to read those entries .
type ConsoleReader struct {
	lines        chan string
	close        func()
	blockEncoder firecore.BlockEncoder

	// ctx   *parseCtx
	done  chan interface{}
	stats *consoleReaderStats

	logger *zap.Logger
	tracer logging.Tracer
}

func NewConsoleReader(lines chan string) (*ConsoleReader, error) {
	l := &ConsoleReader{
		lines: lines,
		close: func() {},
		done:  make(chan interface{}),
	}
	return l, nil
}

// todo: WTF?
func (r *ConsoleReader) Done() <-chan interface{} {
	return r.done
}

func (r *ConsoleReader) Close() {
	r.close()
}

type consoleReaderStats struct {
	lastBlock             bstream.BlockRef
	blockRate             *dmetrics.RateCounter
	blockAverageParseTime *dmetrics.AvgDurationCounter
	transactionRate       *dmetrics.AvgCounter

	cancelPeriodicLogger context.CancelFunc
}

func newConsoleReaderStats() *consoleReaderStats {
	return &consoleReaderStats{
		lastBlock:             bstream.BlockRefEmpty,
		blockRate:             dmetrics.NewPerMinuteLocalRateCounter("blocks"),
		blockAverageParseTime: dmetrics.NewAvgDurationCounter(1*time.Minute, 1*time.Millisecond, "processing block"),
		transactionRate:       dmetrics.NewAvgCounter(1*time.Minute, "trxs"),
	}
}

type parsingStats struct {
	startAt  time.Time
	blockNum uint64
	data     map[string]int
}

func newParsingStats(block uint64) *parsingStats {
	return &parsingStats{
		startAt:  time.Now(),
		blockNum: block,
		data:     map[string]int{},
	}
}

func (s *parsingStats) log() {
	zlog.Info("mindreader block stats",
		zap.Uint64("block_num", s.blockNum),
		zap.Int64("duration", int64(time.Since(s.startAt))),
		zap.Reflect("stats", s.data),
	)
}

func (s *parsingStats) inc(key string) {
	if s == nil {
		return
	}
	k := strings.ToLower(key)
	value := s.data[k]
	value++
	s.data[k] = value
}

// type parseCtx struct {
// 	software     string
// 	majorVersion uint64
// 	minorVersion uint64
// 	// hydrator     antelope.Hydrator

// 	currentBlock *pbbstream.Block
// 	currentTrace *pbantelope.TransactionTrace

// 	abiDecoder     *ABIDecoder
// 	activeBlockNum int64

// 	creationOps       []*creationOp
// 	conversionOptions []antelope.ConversionOption

// 	stats       *parsingStats
// 	globalStats *consoleReaderStats

// 	logger *zap.Logger
// }

func (r *ConsoleReader) ReadBlock() (out *pbbstream.Block, err error) {
	return r.next()
}

const (
	LogPrefix = "DMLOG"
	LogBlock  = "BLOCK"
)

func (r *ConsoleReader) next() (out *pbbstream.Block, err error) {
	for line := range r.lines {
		if !strings.HasPrefix(line, LogPrefix) {
			continue
		}

		tokens := strings.Split(line[len(LogPrefix)+1:], " ")
		if len(tokens) < 2 {
			return nil, fmt.Errorf("invalid log line format: %s", line)
		}

		switch tokens[0] {
		case LogBlock:
			block, err := r.readBlock(tokens[1:])
			if err != nil {
				return nil, fmt.Errorf("read block: %w", err)
			}

			return BlockFromProto(block)

		default:
			if tracer.Enabled() {
				zlog.Debug("skipping unknown deep mind log line", zap.String("line", line))
			}
			continue
		}
	}

	zlog.Info("lines channel has been closed")
	return nil, io.EOF
}

func (r *ConsoleReader) ProcessData(reader io.Reader) error {
	scanner := r.buildScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		r.lines <- line
	}

	if scanner.Err() == nil {
		close(r.lines)
		return io.EOF
	}

	return scanner.Err()
}

func (r *ConsoleReader) buildScanner(reader io.Reader) *bufio.Scanner {
	buf := make([]byte, 50*1024*1024)
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(buf, 50*1024*1024)

	return scanner
}

// Format:
// DMLOG BLOCK <HEIGHT> <ENCODED_BLOCK>
func (r *ConsoleReader) readBlock(params []string) (*pbcodec.Block, error) {
	if err := validateChunk(params, 2); err != nil {
		return nil, fmt.Errorf("invalid log line length: %w", err)
	}

	// <HEIGHT>
	//
	// parse block height
	blockHeight, err := strconv.ParseUint(params[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid block num: %w", err)
	}

	// <ENCODED_BLOCK>
	//
	// hex decode block
	bytes, err := hex.DecodeString(params[1])
	if err != nil {
		return nil, fmt.Errorf("invalid encoded block: %w", err)
	}

	// decode bytes to Block
	block := &pbcodec.Block{}
	err = proto.Unmarshal(bytes, block)
	if err != nil {
		return nil, fmt.Errorf("invalid encoded block: %w", err)
	}

	if blockHeight != block.Height {
		return nil, fmt.Errorf("block height %d from 'DMLOG <height> ...' does not match height %d from block's content", blockHeight, block.Height)
	}

	if tracer.Enabled() {
		zlog.Debug("console reader read block",
			zap.Uint64("height", block.Height),
			zap.Stringer("indep_hash", base64Bytes(block.IndepHash)),
			zap.Stringer("prev_hash", base64Bytes(block.PreviousBlock)),
			zap.Int("trx_count", len(block.Txs)),
		)
	}

	return block, nil
}

func validateChunk(params []string, count int) error {
	if len(params) != count {
		return fmt.Errorf("%d fields required but found %d", count, len(params))
	}
	return nil
}

type base64Bytes []byte

func (b base64Bytes) String() string {
	return base64url.Encode([]byte(b))
}
