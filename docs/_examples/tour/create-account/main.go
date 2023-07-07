package main

import (
	"fmt"

	"github.com/blocto/solana-go-sdk/pkg/hdwallet"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/mr-tron/base58"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	// create a new account
	{
		account := types.NewAccount()
		fmt.Println(account.PublicKey.ToBase58())
		fmt.Println(base58.Encode(account.PrivateKey))
	}

	// from a base58 pirvate key
	{
		account, _ := types.AccountFromBase58("28WJTTqMuurAfz6yqeTrFMXeFd91uzi9i1AW6F5KyHQDS9siXb8TquAuatvLuCEYdggyeiNKLAUr3w7Czmmf2Rav")
		fmt.Println(account.PublicKey.ToBase58())
	}

	// from a private key bytes
	{
		account, _ := types.AccountFromBytes([]byte{
			56, 125, 59, 118, 230, 173, 152, 169, 197, 34,
			168, 187, 217, 160, 119, 204, 124, 69, 52, 136,
			214, 49, 207, 234, 79, 70, 83, 224, 1, 224, 36,
			247, 131, 83, 164, 85, 139, 215, 183, 148, 79,
			198, 74, 93, 156, 157, 208, 99, 221, 127, 51,
			156, 43, 196, 101, 144, 104, 252, 221, 108,
			245, 104, 13, 151,
		})
		fmt.Println(account.PublicKey.ToBase58())
	}

	// from bip 39 (solana cli tool)
	{
		mnemonic := "pill tomorrow foster begin walnut borrow virtual kick shift mutual shoe scatter"
		seed := bip39.NewSeed(mnemonic, "") // (mnemonic, password)
		account, _ := types.AccountFromSeed(seed[:32])
		fmt.Println(account.PublicKey.ToBase58())
	}

	// from bip 44 (phantom)
	{
		mnemonic := "neither lonely flavor argue grass remind eye tag avocado spot unusual intact"
		seed := bip39.NewSeed(mnemonic, "") // (mnemonic, password)
		path := `m/44'/501'/0'/0'`
		derivedKey, _ := hdwallet.Derived(path, seed)
		account, _ := types.AccountFromSeed(derivedKey.PrivateKey)
		fmt.Printf("%v => %v\n", path, account.PublicKey.ToBase58())

		// others
		for i := 1; i < 10; i++ {
			path := fmt.Sprintf(`m/44'/501'/%d'/0'`, i)
			derivedKey, _ := hdwallet.Derived(path, seed)
			account, _ := types.AccountFromSeed(derivedKey.PrivateKey)
			fmt.Printf("%v => %v\n", path, account.PublicKey.ToBase58())
		}
		/*
			m/44'/501'/0'/0' => 5vftMkHL72JaJG6ExQfGAsT2uGVHpRR7oTNUPMs68Y2N
			m/44'/501'/1'/0' => GcXbfQ5yY3uxCyBNDPBbR5FjumHf89E7YHXuULfGDBBv
			m/44'/501'/2'/0' => 7QPgyQwNLqnoSwHEuK8wKy2Y3Ani6EHoZRihTuWkwxbc
			m/44'/501'/3'/0' => 5aE8UprEEWtpVskhxo3f8ETco2kVKiZT9SS3D5Lcg8s2
			m/44'/501'/4'/0' => 5n6afo6LZmzH1J4R38ZCaNSwaztLjd48nWwToLQkCHxp
			m/44'/501'/5'/0' => 2Gr1hWnbaqGXMghicSTHncqV7GVLLddNFJDC7YJoso8M
			m/44'/501'/6'/0' => BNMDY3tCyYbayMzBjZm8RW59unpDWcQRfVmWXCJhLb7D
			m/44'/501'/7'/0' => 9CySTpi4iC85gMW6G4BMoYbNBsdyJrfseHoGmViLha63
			m/44'/501'/8'/0' => ApteF7PmUWS8Lzm6tJPkWgrxSFW5LwYGWCUJ2ByAec91
			m/44'/501'/9'/0' => 6frdqXQAgJMyKwmZxkLYbdGjnYTvUceh6LNhkQt2siQp
		*/
	}
}
