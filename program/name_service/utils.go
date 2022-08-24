package name_service

import (
	"crypto/sha256"

	"github.com/portto/solana-go-sdk/common"
)

var TwitterVerificationAuthority = common.PublicKeyFromString("FvPH7PrVrLGKPfqaf3xJodFTjZriqrAXXLTVWEorTFBi")
var TwitterRootParentRegisteryKey = common.PublicKeyFromString("4YcexoW3r78zz16J2aqmukBLRwGq6rAvWzJpkYAXqebv")
var SolTldAuthority = common.PublicKeyFromString("58PwtjSDuFHuUkYjH9BYnnQKHfwo9reZhC2zMJv9JPkx")

const HashPrefix = "SPL Name Service"

// GetHashName ...
func GetHashName(name string) []byte {
	h := sha256.Sum256([]byte(HashPrefix + name))
	return h[:]
}

// GetNameAccountKey return the pubkey correspond to name
func GetNameAccountKey(hashName []byte, nameClass, nameParent common.PublicKey) common.PublicKey {
	seed := [][]byte{
		hashName,
		nameClass.Bytes(),
		nameParent.Bytes(),
	}
	pubkey, _, _ := common.FindProgramAddress(seed, common.SPLNameServiceProgramID)
	return pubkey
}

// GetTwitterRegistryKey return the pubkey corespond to twitter handle
func GetTwitterRegistryKey(twitterHandle string) common.PublicKey {
	return GetNameAccountKey(
		GetHashName(twitterHandle),
		common.PublicKey{},
		TwitterRootParentRegisteryKey,
	)
}
