# Durable Nonce

A transaction includes a recent blockhash. The recent blockhash will expire after 150 blocks. (arpox. 2 min)
To get rid of it, you can use durable nonce.

## Mechanism

We can trigger the mechanism by

1. use the `nonce` which stored in a nonce account as a recent blockhash
2. make `nonce advance` instruction is the first instruciton
