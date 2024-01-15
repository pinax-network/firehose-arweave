package blockfetcher

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/pinax-network/firehose-arweave/thegarii"
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
	rpcClient                *thegarii.Client
	latest                   uint64
	latestBlockRetryInterval time.Duration
	fetchInterval            time.Duration
	toArwBlock               ToArwBlock
	lastFetchAt              time.Time
	logger                   *zap.Logger
}

func NewBlockFetcher(rpcClient *thegarii.Client, intervalBetweenFetch, latestBlockRetryInterval time.Duration, toArwBlock ToArwBlock, logger *zap.Logger) *BlockFetcher {
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
		latestBlock, err := f.rpcClient.GetCurrentBlock()
		if err != nil {
			return nil, fmt.Errorf("fetching latest block num: %w", err)
		}
		f.latest = latestBlock.Height

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

	rpcBlock, err := f.rpcClient.GetBlockByHeight(blockNum)
	if err != nil {
		return nil, fmt.Errorf("fetching block %d: %w", blockNum, err)
	}

	f.lastFetchAt = time.Now()

	if err != nil {
		return nil, fmt.Errorf("fetching logs for block %d %q: %w", rpcBlock.Height, rpcBlock.Hash, err)
	}

	anyBlock, err := anypb.New(rpcBlock)
	if err != nil {
		return nil, fmt.Errorf("create any block: %w", err)
	}

	return &pbbstream.Block{
		Number:    rpcBlock.Num(),
		Id:        hex.EncodeToString(rpcBlock.IndepHash),
		ParentId:  rpcBlock.PreviousID(),
		Timestamp: timestamppb.New(rpcBlock.GetFirehoseBlockTime()),
		LibNum:    rpcBlock.LIBNum(),
		ParentNum: rpcBlock.GetFirehoseBlockParentNumber(),
		Payload:   anyBlock,
	}, nil
}
