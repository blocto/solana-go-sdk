package rpc

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

func NewRpcClient(endpoint string) RpcClient {
	return RpcClient{endpoint: endpoint}
}

// Call will return body of response. if http code beyond 200~300, the error also returns.
func (c *RpcClient) Call(ctx context.Context, params ...interface{}) ([]byte, error) {
	// prepare payload
	j, err := preparePayload(params)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare payload, err: %v", err)
	}

	// prepare request
	req, err := http.NewRequestWithContext(ctx, "POST", c.endpoint, bytes.NewBuffer(j))
	if err != nil {
		return nil, fmt.Errorf("failed to do http.NewRequestWithContext, err: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	// do request
	httpclient := &http.Client{}
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request, err: %v", err)
	}
	defer res.Body.Close()

	// parse body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body, err: %v", err)
	}

	// check response code
	if res.StatusCode < 200 || res.StatusCode > 300 {
		return body, fmt.Errorf("get status code: %v", res.StatusCode)
	}

	return body, nil
}

func preparePayload(params []interface{}) ([]byte, error) {
	// prepare payload
	j, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  params[0],
		"params":  params[1:],
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

// TODO use Call
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
