package tokenmeta

import (
	"fmt"
	"strings"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/near/borsh-go"
)

const EDITION_MARKER_BIT_SIZE uint64 = 248

type Key borsh.Enum

const (
	KeyUninitialized Key = iota
	KeyEditionV1
	KeyMasterEditionV1
	KeyReservationListV1
	KeyMetadataV1
	KeyReservationListV2
	KeyMasterEditionV2
	KeyEditionMarker
	KeyUseAuthorityRecord
	KeyCollectionAuthorityRecord
)

type Creator struct {
	Address  common.PublicKey
	Verified bool
	Share    uint8
}

type Data struct {
	Name                 string
	Symbol               string
	Uri                  string
	SellerFeeBasisPoints uint16
	Creators             *[]Creator
}

type DataV2 struct {
	Name                 string
	Symbol               string
	Uri                  string
	SellerFeeBasisPoints uint16
	Creators             *[]Creator
	Collection           *Collection
	Uses                 *Uses
}

type metadataPreV11 struct {
	Key                 Key
	UpdateAuthority     common.PublicKey
	Mint                common.PublicKey
	Data                Data
	PrimarySaleHappened bool
	IsMutable           bool
	EditionNonce        *uint8
}

type Metadata struct {
	Key                 Key
	UpdateAuthority     common.PublicKey
	Mint                common.PublicKey
	Data                Data
	PrimarySaleHappened bool
	IsMutable           bool
	EditionNonce        *uint8
	TokenStandard       *TokenStandard
	Collection          *Collection
	Uses                *Uses
	CollectionDetails   *CollectionDetails
	ProgrammableConfig  *ProgrammableConfig
}

type TokenStandard borsh.Enum

const (
	NonFungible TokenStandard = iota
	FungibleAsset
	Fungible
	NonFungibleEdition
	ProgrammableNonFungible
)

type Collection struct {
	Verified bool
	Key      common.PublicKey
}

type Uses struct {
	UseMethod UseMethod
	Remaining uint64
	Total     uint64
}

type UseMethod borsh.Enum

const (
	Burn UseMethod = iota
	Multiple
	Single
)

type CollectionDetails struct {
	Enum borsh.Enum `borsh_enum:"true"`
	V1   CollectionDetailsV1
}

type CollectionDetailsV1 struct {
	Size uint64
}

type ProgrammableConfig struct {
	Enum borsh.Enum `borsh_enum:"true"`
	V1   ProgrammableConfigV1
}

type ProgrammableConfigV1 struct {
	RuleSet *common.PublicKey
}

func MetadataDeserialize(data []byte) (Metadata, error) {
	var metadata Metadata
	err := borsh.Deserialize(&metadata, data)
	if err != nil {
		// https://github.com/samuelvanderwaal/metaboss/issues/121
		// https://github.com/metaplex-foundation/metaplex-program-library/pull/407
		// C.f. https://github.com/metaplex-foundation/metaplex-program-library/blob/master/token-metadata/program/src/deser.rs#L12
		var metadataPreV11 metadataPreV11
		err := borsh.Deserialize(&metadataPreV11, data)
		if err != nil {
			return Metadata{}, fmt.Errorf("failed to deserialize data, err: %v", err)
		} else {
			metadata.Key = metadataPreV11.Key
			metadata.UpdateAuthority = metadataPreV11.UpdateAuthority
			metadata.Mint = metadataPreV11.Mint
			metadata.Data = metadataPreV11.Data
			metadata.PrimarySaleHappened = metadataPreV11.PrimarySaleHappened
			metadata.IsMutable = metadataPreV11.IsMutable
			metadata.EditionNonce = metadataPreV11.EditionNonce
		}
	}
	// trim null byte
	metadata.Data.Name = strings.TrimRight(metadata.Data.Name, "\x00")
	metadata.Data.Symbol = strings.TrimRight(metadata.Data.Symbol, "\x00")
	metadata.Data.Uri = strings.TrimRight(metadata.Data.Uri, "\x00")
	return metadata, nil
}

type MasterEditionV2 struct {
	Key       Key
	Supply    uint64
	MaxSupply *uint64
}
