package sysvar

import (
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/stretchr/testify/assert"
)

func TestDeserializeSlotHashes(t *testing.T) {
	type args struct {
		data  []byte
		owner common.PublicKey
	}
	tests := []struct {
		name string
		args args
		want SlotHashes
		err  error
	}{
		{
			args: args{
				data:  []byte{},
				owner: common.SystemProgramID,
			},
			want: SlotHashes{},
			err:  ErrInvalidAccountOwner,
		},
		{
			args: args{
				data: []byte{
					3, 0, 0, 0, 0, 0, 0, 0,
					3, 0, 0, 0, 0, 0, 0, 0, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
					2, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
					1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				},
				owner: common.SysVarPubkey,
			},
			want: SlotHashes{
				{
					Slot: 3,
					Hash: [32]byte{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
				},
				{
					Slot: 2,
					Hash: [32]byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
				},
				{
					Slot: 1,
					Hash: [32]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				},
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeserializeSlotHashes(tt.args.data, tt.args.owner)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}
