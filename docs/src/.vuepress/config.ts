import { defineUserConfig } from "vuepress";
import type { DefaultThemeOptions } from "vuepress";
import * as path from "path";

export default defineUserConfig<DefaultThemeOptions>({
  lang: "en-US",
  base: "/solana-go-sdk/",
  title: "Solana Development With Go",
  plugins: [["@snippetors/vuepress-plugin-code-copy", { color: "#BBBBBB" }]],
  themeConfig: {
    navbar: [{ text: "GitHub", link: "https://github.com/blocto/solana-go-sdk" }],
    locales: {
      "/": {
        selectLanguageName: "English",
        sidebar: [
          {
            text: "Tour",
            children: [
              {
                text: "Basic",
                children: [
                  { text: "Create Account", link: "/tour/create-account" },
                  { text: "Request Airdrop", link: "/tour/request-airdrop" },
                  { text: "Get Balance", link: "/tour/get-sol-balance" },
                  { text: "Transfer", link: "/tour/transfer" },
                ],
              },
              {
                text: "Token",
                children: [
                  { text: "Create Mint", link: "/tour/create-mint" },
                  { text: "Get Mint", link: "/tour/get-mint" },
                  { text: "Create Token Account", link: "/tour/create-token-account" },
                  { text: "Get Token Account", link: "/tour/get-token-account" },
                  { text: "Mint To", link: "/tour/mint-to" },
                  { text: "Get Balance", link: "/tour/get-token-balance" },
                  { text: "Transfer", link: "/tour/token-transfer" },
                ],
              },
            ],
          },
          {
            text: "NFT",
            children: [
              { text: "Mint a NFT", link: "/nft/mint-a-nft" },
              { text: "Get Metadata", link: "/nft/get-metadata" },
              { text: "Sign Metadata", link: "/nft/sign-metadata" },
            ],
          },
          {
            text: "Advanced",
            children: [
              { text: "Add Memo", link: "/advanced/memo" },
              {
                text: "Durable Nonce",
                link: "/advanced/durable-nonce/README.md",
                children: [
                  { text: "Create Nonce Account", link: "/advanced/durable-nonce/create-nonce-account" },
                  { text: "Get Nonce Account", link: "/advanced/durable-nonce/get-nonce-account" },
                  { text: "Use Nonce", link: "/advanced/durable-nonce/use-nonce" },
                  { text: "Upgrade Nonce", link: "/advanced/durable-nonce/upgrade-nonce" },
                  { text: "Get Nonce Account By Owner", link: "/advanced/durable-nonce/get-nonce-account-by-owner" },
                ],
              },
            ],
          },
          {
            text: "RPC",
            children: [
              { text: "Get Signatures For Address", link: "/rpc/get-signatures-for-address" },
            ],
          },
          {
            text: "Program",
            children: [
              {
                text: "101",
                children: [
                  { text: "Hello", link: "/programs/101/hello" },
                  { text: "Accounts", link: "/programs/101/accounts" },
                  { text: "Data", link: "/programs/101/data" },
                ],
              },
              {
                text: "Stake",
                children: [
                  { text: "Initialize Account", link: "/programs/stake/initialize-account" },
                  { text: "Delegate (stake)", link: "/programs/stake/delegate" },
                  { text: "Deactivate (unstake)", link: "/programs/stake/deactivate" },
                  { text: "Withdraw", link: "/programs/stake/withdraw" },
                ],
              },
            ],
          },
        ],
      },
    },
  },
  markdown: {
    importCode: {
      handleImportPath: (str) => str.replace(/^@/, path.resolve(__dirname, "../../_examples")),
    },
  },
});
