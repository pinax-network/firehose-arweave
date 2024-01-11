package blockfetcher

import (
	"context"
	"time"

	"go.uber.org/zap"

	pbarweave "github.com/pinax-network/firehose-arweave/types/pb/sf/arweave/type/v1"
	pbbstream "github.com/streamingfast/bstream/pb/sf/bstream/v1"
	"github.com/streamingfast/eth-go/rpc"
)

func toArwBlock(in *rpc.Block, receipts map[string]*rpc.TransactionReceipt) (*pbarweave.Block, map[string]bool) {
	// TODO: convert to arweave block
	return nil, nil
}

type ThegariiBlockFetcher struct {
	fetcher *BlockFetcher
}

func NewThegariiBlockFetcher(rpcClient *rpc.Client, intervalBetweenFetch time.Duration, latestBlockRetryInterval time.Duration, logger *zap.Logger) *ThegariiBlockFetcher {
	fetcher := NewBlockFetcher(rpcClient, intervalBetweenFetch, latestBlockRetryInterval, toArwBlock, logger)
	return &ThegariiBlockFetcher{
		fetcher: fetcher,
	}
}

func (f *ThegariiBlockFetcher) PollingInterval() time.Duration { return 1 * time.Second }

func (f *ThegariiBlockFetcher) Fetch(ctx context.Context, blockNum uint64) (*pbbstream.Block, error) {
	return f.fetcher.Fetch(ctx, blockNum)
}
