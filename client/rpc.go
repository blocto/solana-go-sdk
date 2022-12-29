package client

import "github.com/portto/solana-go-sdk/rpc"

func process[A any, B any](fetch func() (rpc.JsonRpcResponse[A], error), convert func(A) (B, error)) (B, error) {
	var output B
	res, err := fetch()
	if err != nil {
		return output, err
	}
	if err = res.GetError(); err != nil {
		return output, err
	}
	return convert(res.GetResult())
}

func value[T any](v rpc.ValueWithContext[T]) (T, error) {
	return v.Value, nil
}

func forward[T any](v T) (T, error) {
	return v, nil
}
