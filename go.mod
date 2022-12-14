module github.com/streamingfast/firehose-arweave

go 1.16

require (
	github.com/ShinyTrinkets/overseer v0.3.0
	github.com/abourget/llerrgroup v0.2.0 // indirect
	github.com/dvsekhvalnov/jose2go v1.5.0
	github.com/golang/protobuf v1.5.2
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.8.1
	github.com/streamingfast/bstream v0.0.2-0.20220916182101-7a027bfdcffb
	github.com/streamingfast/cli v0.0.4-0.20220113202443-f7bcefa38f7e
	github.com/streamingfast/dauth v0.0.0-20220404140613-a40f4cd81626
	github.com/streamingfast/derr v0.0.0-20220526184630-695c21740145
	github.com/streamingfast/dlauncher v0.0.0-20220909121534-7a9aa91dbb32
	github.com/streamingfast/dmetering v0.0.0-20220307162406-37261b4b3de9
	github.com/streamingfast/dmetrics v0.0.0-20220811180000-3e513057d17c
	github.com/streamingfast/dstore v0.1.1-0.20220921155016-7a52fdb3fe5f
	github.com/streamingfast/firehose v0.1.1-0.20220909121738-2f3bc007ea2b
	github.com/streamingfast/firehose-arweave/types v0.0.0-20220509041238-3d3270820c99
	github.com/streamingfast/logging v0.0.0-20220511154537-ce373d264338
	github.com/streamingfast/merger v0.0.3-0.20220909122033-9ca15beb25f5
	github.com/streamingfast/node-manager v0.0.2-0.20220912235129-6c08463b0c01
	github.com/streamingfast/pbgo v0.0.6-0.20220801202203-c32e42ac42a8
	github.com/streamingfast/relayer v0.0.2-0.20220909122435-e67fbc964fd9
	github.com/streamingfast/sf-tools v0.0.0-20220830151952-184d6e9a6bb9
	github.com/stretchr/testify v1.8.0
	go.uber.org/zap v1.21.0
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.0
)

replace github.com/ShinyTrinkets/overseer => github.com/streamingfast/overseer v0.2.1-0.20210326144022-ee491780e3ef

// Pinax dependency replacements for authentication and metering
replace (
	github.com/streamingfast/dauth => github.com/pinax-network/dauth v0.1.0
	github.com/streamingfast/dmetering => github.com/pinax-network/dmetering v0.1.0
	github.com/streamingfast/firehose => github.com/pinax-network/firehose v0.0.0-20220928170757-79bfe0c49eec
)
