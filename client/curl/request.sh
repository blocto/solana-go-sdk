#!/bin/bash

ENDPOINT=https://api.devnet.solana.com

request () {
	echo "==> $1"
	echo "<== $(curl -s "$ENDPOINT" -X POST -H "Content-Type: application/json" -d "$1")"
	echo ""
}
