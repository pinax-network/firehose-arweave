package main

import (
	pbarweave "github.com/pinax-network/firehose-arweave/types/pb/sf/arweave/type/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	firecore "github.com/streamingfast/firehose-core"
	fhCmd "github.com/streamingfast/firehose-core/cmd"
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

func main() {
	fhCmd.Main(Chain())
}

var chain *firecore.Chain[*pbarweave.Block]

func Chain() *firecore.Chain[*pbarweave.Block] {
	if chain != nil {
		return chain
	}

	chain = &firecore.Chain[*pbarweave.Block]{
		ShortName:            "arweave",
		LongName:             "Arweave",
		ExecutableName:       "firearweave",
		FullyQualifiedModule: "github.com/pinax-network/firehose-arweave",
		Version:              version,

		FirstStreamableBlock: 2,

		BlockFactory: func() firecore.Block { return new(pbarweave.Block) },

		RegisterExtraStartFlags: func(flags *pflag.FlagSet) {
			flags.String("reader-node-config-file", "", "Node configuration file, the file is copied inside the {data-dir}/reader/data folder Use {hostname} label to use short hostname in path")
			flags.String("reader-node-genesis-file", "./genesis.json", "Node genesis file, the file is copied inside the {data-dir}/reader/data folder. Use {hostname} label to use short hostname in path")
			flags.String("reader-node-key-file", "./node_key.json", "Node key configuration file, the file is copied inside the {data-dir}/reader/data folder. Use {hostname} label to use with short hostname in path")
			flags.Bool("reader-node-overwrite-node-files", false, "Force download of node-key and config files even if they already exist on the machine.")
		},

		Tools: &firecore.ToolsConfig[*pbarweave.Block]{

			RegisterExtraCmd: func(chain *firecore.Chain[*pbarweave.Block], parent *cobra.Command, zlog *zap.Logger, tracer logging.Tracer) error {
				//toolsCmd.AddCommand(newToolsGenerateNodeKeyCmd(chain))
				//toolsCmd.AddCommand(newToolsBackfillCmd(zlog))
				parent.AddCommand(newPollerCmd(zlog, tracer))
				parent.AddCommand(newSilkwormPollerCmd(zlog, tracer))
				// parent.AddCommand(newCheckBlocksCmd(zlog))

				return nil
			},

			// SanitizeBlockForCompare: sanitizeBlockForCompare,
		},
	}

	return chain
}

// Version value, injected via go build `ldflags` at build time
var version = "dev"
