package nodemanager

import "github.com/streamingfast/logging"

var zlog, _ = logging.PackageLogger("nodemanager", "github.com/streamingfast/firehose-arweave/nodemanager_tests")

func init() {
	logging.InstantiateLoggers()
}
