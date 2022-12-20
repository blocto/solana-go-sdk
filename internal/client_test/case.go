package client_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Param struct {
	Name          string
	RequestBody   string
	ResponseBody  string
	F             func(url string) (any, error)
	ExpectedValue any
	ExpectedError error
}

func TestAll(t *testing.T, params []Param) {
	for _, param := range params {
		t.Run(param.Name, func(t *testing.T) {
			Test(t, param)
		})
	}
}

func Test(t *testing.T, param Param) {
	// setup test server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// check request body match
		body, err := io.ReadAll(req.Body)
		assert.Nil(t, err)
		assert.JSONEq(t, param.RequestBody, string(body))

		// check write response body success
		n, err := rw.Write([]byte(param.ResponseBody))
		assert.Nil(t, err)
		assert.Equal(t, len([]byte(param.ResponseBody)), n)
	}))

	// test function
	got, err := param.F(server.URL)
	assert.Equal(t, param.ExpectedValue, got)
	assert.Equal(t, param.ExpectedError, err)

	server.Close()
}
