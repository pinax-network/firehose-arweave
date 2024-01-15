package thegarii

import (
	// "github.com/goccy/go-json"

	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"

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

// TODO: Handle status code here and return body instead of full response
func (c *Client) get(path string) (*http.Response, error) {
	res, err := http.Get(c.endpoint + path)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func prepareResponse(res *http.Response) (map[string]interface{}, error) {
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var m map[string]interface{}

		// TODO: Redefine Unmarshaler for BigInt handling to simplify code
		err = json.Unmarshal(b, &m)
		if err != nil {
			return nil, err
		}

		return m, nil
	}

	return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
}

func parseBigInt(b interface{}) (*pbarweave.BigInt, error) {
	bi := new(pbarweave.BigInt)

	switch b := b.(type) {
	case string:
		bi.Bytes = []byte(b)
	case float64:
		bi.Bytes = []byte(strconv.FormatFloat(b, 'f', -1, 64))
	default:
		return nil, fmt.Errorf("unexpected type %T", b)
	}

	return bi, nil
}

// TODO: Make it less hardcoded
func (c *Client) parseBlock(res *http.Response) (*pbarweave.Block, error) {
	m, err := prepareResponse(res)
	if err != nil {
		return nil, err
	}

	block := &pbarweave.Block{}

	block.Nonce = []byte(m["nonce"].(string))
	block.PreviousBlock = []byte(m["previous_block"].(string))
	block.Timestamp = uint64(m["timestamp"].(float64))
	block.LastRetarget = uint64(m["last_retarget"].(float64))

	block.Diff, err = parseBigInt(m["diff"])
	if err != nil {
		return nil, err
	}

	block.Height = uint64(m["height"].(float64))
	block.Hash = []byte(m["hash"].(string))
	block.IndepHash = []byte(m["indep_hash"].(string))

	block.Txs = make([]*pbarweave.Transaction, 0)
	txs := m["txs"].([]interface{})

	wg := &sync.WaitGroup{}
	txChan := make(chan *pbarweave.Transaction, len(txs))
	for _, tx := range txs {
		wg.Add(1)

		go func(tx string) {
			defer wg.Done()

			// Get transaction by id
			trx, err := c.GetTxById([]byte(tx))
			if err != nil {
				panic(err)
			}

			// Send transaction to channel
			txChan <- trx
		}(tx.(string))
	}
	wg.Wait()
	close(txChan)

	// Collect transactions from the channel
	// TODO: Should they be in order?
	for trx := range txChan {
		block.Txs = append(block.Txs, trx)
	}

	block.TxRoot = []byte(m["tx_root"].(string))
	block.WalletList = []byte(m["wallet_list"].(string))
	block.RewardAddr = []byte(m["reward_addr"].(string))

	// TODO: Parse Tags

	block.RewardPool, err = parseBigInt(m["reward_pool"])
	if err != nil {
		return nil, err
	}

	block.WeaveSize, err = parseBigInt(m["weave_size"])
	if err != nil {
		return nil, err
	}

	block.BlockSize, err = parseBigInt(m["block_size"])
	if err != nil {
		return nil, err
	}

	if m["cumulative_diff"] != nil {
		block.CumulativeDiff, err = parseBigInt(m["cumulative_diff"])
		if err != nil {
			return nil, err
		}
	}

	if m["hash_list_merkle"] != nil {
		block.HashListMerkle = []byte(m["hash_list_merkle"].(string))
	}

	// TODO: Parse ProofOfAccess

	return block, nil
}

// TODO: Make it less hardcoded
func (c *Client) parseTx(res *http.Response) (*pbarweave.Transaction, error) {
	m, err := prepareResponse(res)
	if err != nil {
		return nil, err
	}

	trx := &pbarweave.Transaction{}

	trx.Format = uint32(m["format"].(float64))
	trx.Id = []byte(m["id"].(string))
	trx.LastTx = []byte(m["last_tx"].(string))
	trx.Owner = []byte(m["owner"].(string))

	// TODO: Parse Tags

	trx.Target = []byte(m["target"].(string))

	trx.Quantity, err = parseBigInt(m["quantity"])
	if err != nil {
		return nil, err
	}

	trx.Data = []byte(m["data"].(string))

	trx.DataSize, err = parseBigInt(m["data_size"])
	if err != nil {
		return nil, err
	}

	trx.DataRoot = []byte(m["data_root"].(string))
	trx.Signature = []byte(m["signature"].(string))

	trx.Reward, err = parseBigInt(m["reward"])
	if err != nil {
		return nil, err
	}

	return trx, nil
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

func (c *Client) GetBlockByHash(hash []byte) (*pbarweave.Block, error) {
	res, err := c.get("block/hash/" + string(hash))
	if err != nil {
		return nil, err
	}

	block, err := c.parseBlock(res)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (c *Client) GetCurrentBlock() (*pbarweave.Block, error) {
	res, err := c.get("current_block")
	if err != nil {
		return nil, err
	}

	block, err := c.parseBlock(res)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (c *Client) GetTxById(id []byte) (*pbarweave.Transaction, error) {
	res, err := c.get("tx/" + string(id))
	if err != nil {
		return nil, err
	}

	tx, err := c.parseTx(res)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// Returns transaction data in base64
func (c *Client) GetTxDataById(id []byte) ([]byte, error) {
	res, err := c.get("tx/" + string(id) + "/data")
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// func (c *Client) GetFirehoseBlockByHeight(height uint64) (*pbarweave.FirehoseBlock, error) {
// 	return nil, nil
// }
