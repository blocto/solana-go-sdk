package tokenmeta

import (
	"fmt"
	"strings"

	"github.com/near/borsh-go"
	"github.com/olegfomenko/solana-go-sdk/common"
)

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

type Metadata struct {
	Key                 Key
	UpdateAuthority     common.PublicKey
	Mint                common.PublicKey
	Data                Data
	PrimarySaleHappened bool
	IsMutable           bool
	EditionNonce        *uint8
}

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
