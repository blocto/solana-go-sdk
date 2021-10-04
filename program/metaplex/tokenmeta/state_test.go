package tokenmeta

import (
	"github.com/olegfomenko/solana-go-sdk/common/pointer"
	"testing"

	"github.com/olegfomenko/solana-go-sdk/common"
	"github.com/stretchr/testify/assert"
)

func TestMetadataDeserialize(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want Metadata
		err  error
	}{
		{
			args: args{
				data: []byte{0x4, 0xb5, 0x1f, 0xa2, 0x9b, 0x1d, 0xdb, 0xc2, 0x64, 0x1f, 0x10, 0xf9, 0xd3, 0xc9, 0x21, 0x39, 0x5a, 0x89, 0xb3, 0x44, 0x4, 0xb4, 0x98, 0xa5, 0xe9, 0x89, 0xd0, 0xc7, 0xa4, 0xe7, 0xb1, 0xf8, 0xc2, 0xeb, 0x17, 0x1a, 0x88, 0x2d, 0x89, 0x5e, 0x4e, 0xe6, 0x5a, 0x27, 0x14, 0xf, 0xd5, 0x76, 0x93, 0xd3, 0x6f, 0x8d, 0x73, 0x54, 0x6f, 0x18, 0xc6, 0x8b, 0x8c, 0xec, 0xe8, 0xc7, 0x8c, 0xfb, 0xef, 0x20, 0x0, 0x0, 0x0, 0x44, 0x65, 0x67, 0x65, 0x6e, 0x20, 0x41, 0x70, 0x65, 0x20, 0x23, 0x31, 0x38, 0x32, 0x39, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x44, 0x41, 0x50, 0x45, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc8, 0x0, 0x0, 0x0, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x72, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x36, 0x65, 0x50, 0x63, 0x77, 0x33, 0x67, 0x32, 0x77, 0x49, 0x2d, 0x6b, 0x4a, 0x52, 0x46, 0x66, 0x74, 0x55, 0x5f, 0x64, 0x51, 0x44, 0x61, 0x50, 0x37, 0x30, 0x6d, 0x79, 0x4b, 0x70, 0x4d, 0x78, 0x53, 0x56, 0x45, 0x44, 0x64, 0x47, 0x6c, 0x62, 0x4e, 0x65, 0x30, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa4, 0x1, 0x1, 0x5, 0x0, 0x0, 0x0, 0x79, 0x82, 0x3c, 0xdb, 0x6c, 0xe5, 0x8, 0x73, 0xdb, 0x3f, 0x89, 0x5e, 0x24, 0x3f, 0x94, 0xd8, 0x22, 0xb, 0x98, 0x38, 0x70, 0x2e, 0xf9, 0x7b, 0xba, 0xeb, 0x69, 0x70, 0xdd, 0x23, 0x94, 0xff, 0x0, 0x27, 0x7a, 0x97, 0xa8, 0xdd, 0x89, 0x92, 0x55, 0x7b, 0xb1, 0x77, 0x56, 0x8, 0x18, 0xd, 0x5, 0xc7, 0xf3, 0x41, 0x86, 0x5a, 0xdd, 0x90, 0xde, 0xf0, 0x8a, 0x19, 0x69, 0xf3, 0x40, 0x1, 0x6c, 0xd6, 0x0, 0x19, 0xf3, 0x2d, 0xb6, 0x39, 0x4e, 0xf7, 0x48, 0x97, 0x6b, 0x14, 0xec, 0xfb, 0xc2, 0x76, 0xda, 0xc4, 0x2f, 0x8, 0x17, 0x61, 0x62, 0x18, 0x4b, 0x4b, 0xda, 0x98, 0x3, 0x6, 0x20, 0x6d, 0xd5, 0x4a, 0x0, 0x19, 0x5c, 0xfd, 0x4a, 0x8f, 0xee, 0xc6, 0x1e, 0x42, 0x28, 0xf0, 0x28, 0x6c, 0x5d, 0x2b, 0x0, 0x51, 0xd, 0x26, 0x15, 0x26, 0x3a, 0x20, 0xed, 0x58, 0x3f, 0x19, 0xda, 0x11, 0x27, 0x44, 0xed, 0x10, 0x0, 0xa, 0xb5, 0x1f, 0xa2, 0x9b, 0x1d, 0xdb, 0xc2, 0x64, 0x1f, 0x10, 0xf9, 0xd3, 0xc9, 0x21, 0x39, 0x5a, 0x89, 0xb3, 0x44, 0x4, 0xb4, 0x98, 0xa5, 0xe9, 0x89, 0xd0, 0xc7, 0xa4, 0xe7, 0xb1, 0xf8, 0xc2, 0x1, 0x1, 0x1, 0x0, 0x1, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			},
			want: Metadata{
				Key:             4,
				UpdateAuthority: common.PublicKeyFromString("DC2mkgwhy56w3viNtHDjJQmc7SGu2QX785bS4aexojwX"),
				Mint:            common.PublicKeyFromString("GphF2vTuzhwhLWBWWvD8y5QLCPp1aQC5EnzrWsnbiWPx"),
				Data: Data{
					Name:                 "Degen Ape #1829",
					Symbol:               "DAPE",
					Uri:                  "https://arweave.net/6ePcw3g2wI-kJRFftU_dQDaP70myKpMxSVEDdGlbNe0",
					SellerFeeBasisPoints: 420,
					Creators: &[]Creator{
						{
							Address:  common.PublicKeyFromString("9BKWqDHfHZh9j39xakYVMdr6hXmCLHH5VfCpeq2idU9L"),
							Verified: false,
							Share:    39,
						},
						{
							Address:  common.PublicKeyFromString("9FYsKrNuEweb55Wa2jaj8wTKYDBvuCG3huhakEj96iN9"),
							Verified: false,
							Share:    25,
						},
						{
							Address:  common.PublicKeyFromString("HNGVuL5kqjDehw7KR63w9gxow32sX6xzRNgLb8GkbwCM"),
							Verified: false,
							Share:    25,
						},
						{
							Address:  common.PublicKeyFromString("7FzXBBPjzrNJbm9MrZKZcyvP3ojVeYPUG2XkBPVZvuBu"),
							Verified: false,
							Share:    10,
						},
						{
							Address:  common.PublicKeyFromString("DC2mkgwhy56w3viNtHDjJQmc7SGu2QX785bS4aexojwX"),
							Verified: true,
							Share:    1,
						},
					},
				},
				PrimarySaleHappened: true,
				IsMutable:           false,
				EditionNonce:        pointer.Uint8(255),
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MetadataDeserialize(tt.args.data)
			assert.Equal(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}
