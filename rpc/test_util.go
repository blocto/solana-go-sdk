package rpc

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testRpcCallParam struct {
	RequestBody      string
	ResponseBody     string
	RpcCall          func(RpcClient) (interface{}, error)
	ExpectedResponse interface{}
	ExpectedError    error
}

func testRpcCall(t *testing.T, param testRpcCallParam) {
	// setup test server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		assert.Nil(t, err)
		assert.JSONEq(t, param.RequestBody, string(body))
		n, err := rw.Write([]byte(param.ResponseBody))
		assert.Nil(t, err)
		assert.Equal(t, len([]byte(param.ResponseBody)), n)
	}))

	// test call
	got, err := param.RpcCall(NewRpcClient(server.URL))
	assert.Equal(t, param.ExpectedError, err)
	assert.Equal(t, param.ExpectedResponse, got)

	server.Close()
}
