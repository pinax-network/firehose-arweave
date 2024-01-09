package blockfetcher

import (
	"context"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"github.com/abourget/llerrgroup"
	pbarweave "github.com/pinax-network/firehose-arweave/types/pb/sf/arweave/type/v1"
	pbbstream "github.com/streamingfast/bstream/pb/sf/bstream/v1"

	"github.com/streamingfast/eth-go/rpc"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const CONFIRMS uint64 = 20

type ToArwBlock func(in *rpc.Block, receipts map[string]*rpc.TransactionReceipt) (*pbarweave.Block, map[string]bool)

type BlockFetcher struct {
	rpcClient                *rpc.Client
	latest                   uint64
	latestBlockRetryInterval time.Duration
	fetchInterval            time.Duration
	toArwBlock               ToArwBlock
	lastFetchAt              time.Time
	logger                   *zap.Logger
}

func NewBlockFetcher(rpcClient *rpc.Client, intervalBetweenFetch, latestBlockRetryInterval time.Duration, toArwBlock ToArwBlock, logger *zap.Logger) *BlockFetcher {
	return &BlockFetcher{
		rpcClient:                rpcClient,
		latestBlockRetryInterval: latestBlockRetryInterval,
		toArwBlock:               toArwBlock,
		fetchInterval:            intervalBetweenFetch,
		logger:                   logger,
	}
}

func (f *BlockFetcher) Fetch(ctx context.Context, blockNum uint64) (block *pbbstream.Block, err error) {
	f.logger.Debug("fetching block", zap.Uint64("block_num", blockNum))
	for f.latest < blockNum {
		f.latest, err = f.rpcClient.LatestBlockNum(ctx)
		if err != nil {
			return nil, fmt.Errorf("fetching latest block num: %w", err)
		}

		f.logger.Info("got latest block", zap.Uint64("latest", f.latest), zap.Uint64("block_num", blockNum))

		if f.latest < blockNum {
			time.Sleep(f.latestBlockRetryInterval)
			continue
		}
		break
	}

	sinceLastFetch := time.Since(f.lastFetchAt)
	if sinceLastFetch < f.fetchInterval {
		time.Sleep(f.fetchInterval - sinceLastFetch)
	}

	rpcBlock, err := f.rpcClient.GetBlockByNumber(ctx, rpc.BlockNumber(blockNum), rpc.WithGetBlockFullTransaction())
	if err != nil {
		return nil, fmt.Errorf("fetching block %d: %w", blockNum, err)
	}

	receipts, err := FetchReceipts(ctx, rpcBlock, f.rpcClient)
	if err != nil {
		return nil, fmt.Errorf("fetching receipts for block %d %q: %w", rpcBlock.Number, rpcBlock.Hash.Pretty(), err)
	}

	f.logger.Debug("fetched receipts", zap.Int("count", len(receipts)))

	f.lastFetchAt = time.Now()

	if err != nil {
		return nil, fmt.Errorf("fetching logs for block %d %q: %w", rpcBlock.Number, rpcBlock.Hash.Pretty(), err)
	}

	arwBlock, _ := f.toArwBlock(rpcBlock, receipts)
	anyBlock, err := anypb.New(arwBlock)
	if err != nil {
		return nil, fmt.Errorf("create any block: %w", err)
	}

	return &pbbstream.Block{
		Number:    arwBlock.Num(),
		Id:        hex.EncodeToString(arwBlock.IndepHash),
		ParentId:  arwBlock.PreviousID(),
		Timestamp: timestamppb.New(arwBlock.GetFirehoseBlockTime()),
		LibNum:    arwBlock.LIBNum(),
		ParentNum: arwBlock.GetFirehoseBlockParentNumber(),
		Payload:   anyBlock,
	}, nil
}

func FetchReceipts(ctx context.Context, block *rpc.Block, client *rpc.Client) (out map[string]*rpc.TransactionReceipt, err error) {
	out = make(map[string]*rpc.TransactionReceipt)
	lock := sync.Mutex{}

	eg := llerrgroup.New(10)
	for _, tx := range block.Transactions.Transactions {
		if eg.Stop() {
			continue // short-circuit the loop if we got an error
		}
		eg.Go(func() error {
			receipt, err := client.TransactionReceipt(ctx, tx.Hash)
			if err != nil {
				return fmt.Errorf("fetching receipt for tx %q: %w", tx.Hash.Pretty(), err)
			}
			lock.Lock()
			out[tx.Hash.Pretty()] = receipt
			lock.Unlock()
			return err
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return
}
