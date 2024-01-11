package main

import (
	"path"

	"github.com/pinax-network/firehose-arweave/thegarii"
	"github.com/spf13/cobra"
	"github.com/streamingfast/cli/sflags"
	firecore "github.com/streamingfast/firehose-core"
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

func newPollerCmd(logger *zap.Logger, tracer logging.Tracer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "poller",
		Short: "poll blocks from different sources",
	}

	cmd.AddCommand(newThegariiPollerCmd(logger, tracer))
	return cmd
}

func newThegariiPollerCmd(logger *zap.Logger, tracer logging.Tracer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "thegarii <rpc-endpoint> <first-streamable-block>",
		Short: "poll blocks from thegarii rpc",
		Args:  cobra.ExactArgs(2),
		RunE:  pollerRunE(logger, tracer),
	}
	cmd.Flags().Duration("interval-between-fetch", 0, "interval between fetch")

	return cmd
}

func pollerRunE(logger *zap.Logger, tracer logging.Tracer) firecore.CommandExecutor {
	return func(cmd *cobra.Command, args []string) (err error) {
		// ctx := cmd.Context()

		rpcEndpoint := args[0]

		dataDir := sflags.MustGetString(cmd, "data-dir")
		stateDir := path.Join(dataDir, "poller-state")

		logger.Info("launching firehose-arweave poller", zap.String("rpc_endpoint", rpcEndpoint), zap.String("data_dir", dataDir), zap.String("state_dir", stateDir))

		// rpcClient := rpc.NewClient(rpcEndpoint)

		// firstStreamableBlock, err := strconv.ParseUint(args[1], 10, 64)
		// if err != nil {
		// 	return fmt.Errorf("unable to parse first streamable block %d: %w", firstStreamableBlock, err)
		// }

		// fetchInterval := sflags.MustGetDuration(cmd, "interval-between-fetch")

		// fetcher := blockfetcher.NewThegariiBlockFetcher(rpcClient, fetchInterval, 1*time.Second, logger)
		// handler := blockpoller.NewFireBlockHandler("type.googleapis.com/sf.ethereum.type.v2.Block")
		// poller := blockpoller.New(fetcher, handler, blockpoller.WithStoringState(stateDir), blockpoller.WithLogger(logger))

		// latestBlock, err := rpcClient.GetBlockByNumber(ctx, rpc.LatestBlock)
		// if err != nil {
		// 	return fmt.Errorf("getting latest block: %w", err)
		// }

		// latestFinalizedBlock, err := rpcClient.GetBlockByNumber(ctx, rpc.BlockNumber(uint64(latestBlock.Number)))
		// if err != nil {
		// 	return fmt.Errorf("getting latest finalized block: %w", err)
		// }

		// err = poller.Run(ctx, firstStreamableBlock, bstream.NewBlockRef(latestFinalizedBlock.Hash.String(), uint64(latestFinalizedBlock.Number)))
		// if err != nil {
		// 	return fmt.Errorf("running poller: %w", err)
		// }

		rpcClient := thegarii.NewClient(rpcEndpoint)
		_, err = rpcClient.GetBlockByHeight(200)
		if err != nil {
			return err
		}

		return nil
	}
}
