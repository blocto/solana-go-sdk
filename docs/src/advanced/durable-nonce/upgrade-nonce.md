# Upgrade Nonce

Due to [2022/06/01 Solana outage](https://solana.com/news/06-01-22-solana-mainnet-beta-outage-report-2). All nonce account should be upgraded. You can update it either

1. Advance Nonce (need origin authority signature)

@[code](@/advanced/durable-nonce/advance-nonce/main.go)

2. Upgrade Nonce (you can use a random fee payer)

@[code](@/advanced/durable-nonce/upgrade-nonce-account/main.go)
