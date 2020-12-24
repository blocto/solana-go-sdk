package client

import (
	"bytes"
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	DevnetRPCEndpoint  = "https://devnet.solana.com"
	TestnetRPCEndpoint = "https://testnet.solana.com"
	MainnetRPCEndpoint = "https://api.mainnet-beta.solana.com"
)

type Commitment string

const (
	CommitmentMax          Commitment = "max"
	CommitmentRoot                    = "root"
	CommitmentSingleGossip            = "singleGossip"
	CommitmentRecnet                  = "recent"
)

type Client struct {
	endpoint string
}

func NewClient(endpoint string) *Client {
	return &Client{endpoint: endpoint}
}

func (s *Client) request(method string, params []interface{}, response interface{}) error {
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
	req, err := http.NewRequest("POST", s.endpoint, bytes.NewBuffer(j))
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

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Context struct {
	Slot uint64 `json:"slot"`
}

type GeneralResponse struct {
	JsonRPC string        `json:"jsonrpc"`
	ID      uint64        `json:"id"`
	Error   ErrorResponse `json:"error"`
}

func (s *Client) GenerateKey() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	return ed25519.GenerateKey(nil)
}
