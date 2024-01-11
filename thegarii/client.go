package thegarii

import (
	// "github.com/goccy/go-json"

	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	pbarweave "github.com/pinax-network/firehose-arweave/types/pb/sf/arweave/type/v1"
)

const DEFAULT_ENDPOINT = "https://arweave.net/"

type Client struct {
	client   *http.Client
	endpoint string
}

func NewClient(endpoint string) *Client {
	if endpoint == "" {
		endpoint = DEFAULT_ENDPOINT
	}

	if endpoint[len(endpoint)-1] != '/' {
		endpoint += "/"
	}

	return &Client{
		client:   http.DefaultClient,
		endpoint: endpoint,
	}
}

func (c *Client) get(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.endpoint+path, nil)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func (c *Client) parseBlock(res *http.Response) (*pbarweave.Block, error) {
	// defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var m map[string]interface{}
		err = json.Unmarshal(b, &m)
		if err != nil {
			return nil, err
		}

		block := &pbarweave.Block{}

		block.Nonce = []byte(m["nonce"].(string))
		block.PreviousBlock = []byte(m["previous_block"].(string))
		block.Timestamp = uint64(m["timestamp"].(float64))
		block.LastRetarget = uint64(m["last_retarget"].(float64))
		// block.Diff = uint64(m["diff"].()["bytes"].(float64))
		block.Height = uint64(m["height"].(float64))
		block.Hash = []byte(m["hash"].(string))
		// block.Txs = []byte(m["txs"].(string))
		block.TxRoot = []byte(m["tx_root"].(string))

		return nil, nil
	}

	return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
}

func (c *Client) GetBlockByHeight(height uint64) (*pbarweave.Block, error) {
	res, err := c.get("block/height/" + strconv.FormatUint(height, 10))
	if err != nil {
		return nil, err
	}

	block, err := c.parseBlock(res)
	if err != nil {
		return nil, err
	}

	return block, nil
}
