package token_metadata

import (
	"strconv"

	"github.com/blocto/solana-go-sdk/common"
)

func GetTokenMetaPubkey(mint common.PublicKey) (common.PublicKey, error) {
	metadataAccount, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			common.MetaplexTokenMetaProgramID.Bytes(),
			mint.Bytes(),
		},
		common.MetaplexTokenMetaProgramID,
	)
	if err != nil {
		return common.PublicKey{}, err
	}
	return metadataAccount, nil
}

func GetMasterEdition(mint common.PublicKey) (common.PublicKey, error) {
	msaterEdtion, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			common.MetaplexTokenMetaProgramID.Bytes(),
			mint.Bytes(),
			[]byte("edition"),
		},
		common.MetaplexTokenMetaProgramID,
	)
	if err != nil {
		return common.PublicKey{}, err
	}
	return msaterEdtion, nil
}

func GetEditionMark(mint common.PublicKey, edition uint64) (common.PublicKey, error) {
	editionNumber := edition / EDITION_MARKER_BIT_SIZE
	pubkey, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			common.MetaplexTokenMetaProgramID.Bytes(),
			mint.Bytes(),
			[]byte("edition"),
			[]byte(strconv.FormatUint(editionNumber, 10)),
		},
		common.MetaplexTokenMetaProgramID,
	)
	return pubkey, err
}
