package tokenmeta

import (
	"fmt"
	"strings"

	"github.com/near/borsh-go"
	"github.com/portto/solana-go-sdk/common"
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

type Metadata struct {
	Key                 Key
	UpdateAuthority     common.PublicKey
	Mint                common.PublicKey
	Data                Data
	PrimarySaleHappened bool
	IsMutable           bool
	EditionNonce        *uint8
}

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

func MetadataDeserialize(data []byte) (Metadata, error) {
	var metadata Metadata
	err := borsh.Deserialize(&metadata, data)
	if err != nil {
		return Metadata{}, fmt.Errorf("failed to deserialize data, err: %v", err)
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
