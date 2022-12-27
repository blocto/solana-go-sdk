package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	LocalnetRPCEndpoint = "http://localhost:8899"
	DevnetRPCEndpoint   = "https://api.devnet.solana.com"
	TestnetRPCEndpoint  = "https://api.testnet.solana.com"
	MainnetRPCEndpoint  = "https://api.mainnet-beta.solana.com"
)

type JsonRpcRequest struct {
	JsonRpc string `json:"jsonrpc"`
	Id      uint64 `json:"id"`
	Method  string `json:"method"`
	Params  []any  `json:"params,omitempty"`
}

type JsonRpcResponse[T any] struct {
	JsonRpc string        `json:"jsonrpc"`
	Id      uint64        `json:"id"`
	Result  T             `json:"result"`
	Error   *JsonRpcError `json:"error,omitempty"`
}

type JsonRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (e *JsonRpcError) Error() string {
	s, err := json.Marshal(e)
	if err == nil {
		return string(s)
	}

	// ideally, it should never reach here
	return fmt.Sprintf("failed to marshal JsonRpcError, err: %v, code: %v, message: %v, data: %v", err, e.Code, e.Message, e.Data)
}

type ValueWithContext[T any] struct {
	Context Context `json:"context"`
	Value   T       `json:"value"`
}

type RpcClient struct {
	endpoint   string
	httpClient *http.Client
}

func NewRpcClient(endpoint string) RpcClient { return New(WithEndpoint(endpoint)) }

// New applies the given options to the rpc client being created. if no options
// is passed, it defaults to a bare bone http client and solana mainnet
func New(opts ...Option) RpcClient {

	client := &RpcClient{}

	setDefaultOptions(client)

	for _, opt := range opts {
		opt(client)
	}

	return *client
}

// Call will return body of response. if http code beyond 200~300, the error also returns.
func (c *RpcClient) Call(ctx context.Context, params ...any) ([]byte, error) {
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
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request, err: %v", err)
	}
	defer res.Body.Close()

	// parse body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body, err: %v", err)
	}

	// check response code
	if res.StatusCode < 200 || res.StatusCode > 300 {
		return body, fmt.Errorf("get status code: %v", res.StatusCode)
	}

	return body, nil
}

func preparePayload(params []any) ([]byte, error) {
	// prepare payload
	j, err := json.Marshal(JsonRpcRequest{
		JsonRpc: "2.0",
		Id:      1,
		Method:  params[0].(string),
		Params:  params[1:],
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func call[T any](c *RpcClient, ctx context.Context, params ...any) (T, error) {
	var output T

	// rpc call
	body, err := c.Call(ctx, params...)
	if err != nil {
		return output, fmt.Errorf("rpc: call error, err: %v, body: %v", err, string(body))
	}

	// transfer data
	err = json.Unmarshal(body, &output)
	if err != nil {
		return output, fmt.Errorf("rpc: failed to json decode body, err: %v", err)
	}

	return output, nil
}
