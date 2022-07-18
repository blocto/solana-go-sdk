package rpc

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestOption_WithHTTPClient(t *testing.T) {

	h := &http.Client{
		Timeout: time.Minute * 20,
	}

	c := New(WithHTTPClient(h))

	require.Equal(t, h, c.httpClient)
}

func TestOption_WithEndpoint(t *testing.T) {

	endpoint := DevnetRPCEndpoint

	c := New(WithEndpoint(endpoint))

	require.Equal(t, endpoint, c.endpoint)
}
