package main

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/rpc"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	target := "Memo1UhkJRfHyvLMcVucJwxXeuD728EqVDDwQDxFMNo"

	// get all (limit is between 1 ~ 1,000, default is 1,000)
	{
		res, err := c.GetSignaturesForAddress(context.Background(), target)
		if err != nil {
			log.Fatalf("failed to GetSignaturesForAddress, err: %v", err)
		}
		spew.Dump(res)
	}

	// get latest X tx
	{
		res, err := c.GetSignaturesForAddressWithConfig(
			context.Background(),
			target,
			rpc.GetSignaturesForAddressConfig{
				Limit: 5,
			},
		)
		if err != nil {
			log.Fatalf("failed to GetSignaturesForAddress, err: %v", err)
		}
		spew.Dump(res)
	}

	/*
		if a txhash list like:

		(new)
		3gwJwVorVprqZmm1ULAp9bKQy6sQ7skG11XYdMSQigA936MBANcSBy6NcNJF2yPf9ycgzZ6vFd4pjAY7qSko61Au
		wTEnw3vpBthzLUD6gv9B3aC4dQNp4hews85ipM3w9MGZAh38HZ2im9LaWY7aRusVN5Wj33mNvqSRDNyC43u6GQs
		3e6dRv5KnvpU43VjVjbsubvPR1yFK9b922WcTugyTBSdWdToeCK16NccSaxY6XJ5yi51UswP3ZDe3VJBZTVg2MCW
		2nYnHvbVuwmYeara3VjoCt9uS8ZXrSra5DRK7QBT8i5acoBiSK3FQY2vsaDSJQ6QX5i1pkvyRRjL1oUATMLZEsqy
		2uFaNDgQWZsgZvR6s3WQKwaCxFgS4ML7xrZyAqgmuTSEuGmrWyCcTrjtajr6baYR6FaVLZ4PWgyt55EmTcT8S7Sg
		4XGVHHpLW99AUFEd6RivasG57vqu4EMMNdcQdmphepmW484dMYtWLkYw4nSNnSpKiDoYDbSu9ksxECNKBk2JEyHQ
		3kjLJokcYqAhQjERCVutv5gdUuQ1HsxSCcFsJdQbqNkqd5ML8WRaZJguZgpWH8isCfyEN8YktxxPPNJURhAtvUKE
		(old)
	*/

	// you can fetch the last 3 tx by
	{
		res, err := c.GetSignaturesForAddressWithConfig(
			context.Background(),
			target,
			rpc.GetSignaturesForAddressConfig{
				Before: "2nYnHvbVuwmYeara3VjoCt9uS8ZXrSra5DRK7QBT8i5acoBiSK3FQY2vsaDSJQ6QX5i1pkvyRRjL1oUATMLZEsqy",
				Limit:  3,
			},
		)
		if err != nil {
			log.Fatalf("failed to GetSignaturesForAddress, err: %v", err)
		}
		spew.Dump(res)
	}

	// you can fetch the latest 3 tx by `until`
	// * the result will be different if there are some newer txs added.
	{
		res, err := c.GetSignaturesForAddressWithConfig(
			context.Background(),
			target,
			rpc.GetSignaturesForAddressConfig{
				Until: "2nYnHvbVuwmYeara3VjoCt9uS8ZXrSra5DRK7QBT8i5acoBiSK3FQY2vsaDSJQ6QX5i1pkvyRRjL1oUATMLZEsqy",
				Limit: 3,
			},
		)
		if err != nil {
			log.Fatalf("failed to GetSignaturesForAddress, err: %v", err)
		}
		spew.Dump(res)
	}
}
