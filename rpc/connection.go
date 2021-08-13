package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	DevnetRPCEndpoint  = "https://api.devnet.solana.com"
	TestnetRPCEndpoint = "https://api.testnet.solana.com"
	MainnetRPCEndpoint = "https://api.mainnet-beta.solana.com"
)

type RpcClient struct {
	endpoint string
}

func NewClient(endpoint string) *RpcClient {
	return &RpcClient{endpoint: endpoint}
}

func (s *RpcClient) request(ctx context.Context, method string, params []interface{}, response interface{}) error {
	// post data
	j, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      0,
		"method":  method,
		"params":  params,
	})
	if err != nil {
		return err
	}

	// post request
	req, err := http.NewRequestWithContext(ctx, "POST", s.endpoint, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	// http client and send request
	httpclient := &http.Client{}
	res, err := httpclient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// parse body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(body) != 0 {
		if err := json.Unmarshal(body, &response); err != nil {
			return err
		}
	}

	// return result
	if res.StatusCode < 200 || res.StatusCode > 300 {
		return fmt.Errorf("get status code: %d", res.StatusCode)
	}
	return nil
}
