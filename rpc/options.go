package rpc

import (
	"net/http"
)

// Option is a configuration type for the Client
type Option func(RpcClient)

// HTTPClient is an Option type that allows you provide your own HTTP client
func WithHTTPClient(h *http.Client) Option {
	return func(r RpcClient) {
		r.httpClient = h
	}
}

// WithEndpoint is an Option that allows you configure the rpc endpoint that our
// client will point to
func WithEndpoint(endpoint string) Option {
	return func(r RpcClient) {
		r.endpoint = endpoint
	}
}

func setDefaultOptions(r RpcClient) {
	r.httpClient = &http.Client{}
	r.endpoint = MainnetRPCEndpoint
}
