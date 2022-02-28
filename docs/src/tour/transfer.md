# Transfer

We will build a transaction to transfer our SOL out.

A quick guide for a transaction is that:

1. a transactions is composed by some signatures + a message
2. a message is composed by one or more instructions + a blockhash

## Full Code

@[code](@/tour/transfer-sol/main.go)

::: tip

1. `fee payer` and `from` can be different wallets but both of them need to sign the transaction
2. a transaction can composed by many instructions so you can do something like A => B, B => C, C => A and D is the fee payer.
:::
