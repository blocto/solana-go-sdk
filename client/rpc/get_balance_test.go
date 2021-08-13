package rpc

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetBalance(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		ctx        context.Context
		base58Addr string
		cfg        GetBalanceConfig
	}
	tests := []struct {
		name                string
		expectedRequestBody string
		responseBody        string
		args                args
		want                GetBalanceResponse
		wantErr             bool
	}{
		{
			args: args{
				ctx:        context.TODO(),
				base58Addr: "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
				cfg:        GetBalanceConfig{},
			},
			expectedRequestBody: `{"jsonrpc":"2.0", "id":1, "method":"getBalance", "params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"]}`,
			responseBody:        `{"jsonrpc":"2.0","result":{"context":{"slot":73914708},"value":6999995000},"id":1}`,
			want: GetBalanceResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetBalanceResult{
					Context: Context{
						Slot: 73914708,
					},
					Value: 6999995000,
				},
			},
			wantErr: false,
		},
		{
			args: args{
				ctx:        context.TODO(),
				base58Addr: "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
				cfg:        GetBalanceConfig{Commitment: CommitmentFinalized},
			},
			expectedRequestBody: `{"jsonrpc":"2.0", "id":1, "method":"getBalance", "params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7", {"commitment": "finalized"}]}`,
			responseBody:        `{"jsonrpc":"2.0","result":{"context":{"slot":73914708},"value":6999995000},"id":1}`,
			want: GetBalanceResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetBalanceResult{
					Context: Context{
						Slot: 73914708,
					},
					Value: 6999995000,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				body, err := ioutil.ReadAll(req.Body)

				if err != nil {
					t.Errorf("read request body error")
					return
				}
				assert.JSONEq(tt.expectedRequestBody, string(body))
				rw.Write([]byte(tt.responseBody))
			}))
			c := NewRpcClient(server.URL)
			got, err := c.GetBalance(tt.args.ctx, tt.args.base58Addr, tt.args.cfg)

			if (err != nil) != tt.wantErr {

				t.Errorf("Client.GetBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetBalance() = %v, want %v", got, tt.want)
			}
			server.Close()
		})
	}
}
