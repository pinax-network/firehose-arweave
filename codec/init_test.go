package codec

import (
	"github.com/pinax-network/firehose-arweave/types"
	"github.com/streamingfast/logging"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var zlogTest, _ = logging.PackageLogger("firearweave", "github.com/pinax-network/firehose-arweave/codec.tests")

func init() {
	types.InitFireCore()
	logging.InstantiateLoggers()
}

type ObjectReader func() (interface{}, error)

func MarshalIndentToString(m proto.Message, indent string) (string, error) {
	res, err := protojson.MarshalOptions{Indent: indent}.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(res), err
}
