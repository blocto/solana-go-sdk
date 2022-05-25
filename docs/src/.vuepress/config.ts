import { defineUserConfig } from "vuepress";
import type { DefaultThemeOptions } from "vuepress";
import * as path from "path";

export default defineUserConfig<DefaultThemeOptions>({
  lang: "en-US",
  base: "/solana-go-sdk/",
  title: "Solana Development With Go",
  plugins: [["@snippetors/vuepress-plugin-code-copy", { color: "#BBBBBB" }]],
  themeConfig: {
    navbar: [{ text: "GitHub", link: "https://github.com/OldSmokeGun/solana-go-sdk" }],
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
