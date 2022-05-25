package cmptbdgprog

import (
	"reflect"
	"testing"

	"github.com/OldSmokeGun/solana-go-sdk/common"
	"github.com/OldSmokeGun/solana-go-sdk/types"
)

func TestRequestUnits(t *testing.T) {
	type args struct {
		param RequestUnitsParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: RequestUnitsParam{
					Units:         1000,
					AdditionalFee: 2000,
				},
			},
			want: types.Instruction{
				ProgramID: common.ComputeBudgetProgramID,
				Accounts:  []types.AccountMeta{},
				Data:      []byte{0, 232, 3, 0, 0, 208, 7, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RequestUnits(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestUnits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestHeapFrame(t *testing.T) {
	type args struct {
		param RequestHeapFrameParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: RequestHeapFrameParam{
					Bytes: 1000,
				},
			},
			want: types.Instruction{
				ProgramID: common.ComputeBudgetProgramID,
				Accounts:  []types.AccountMeta{},
				Data:      []byte{1, 232, 3, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RequestHeapFrame(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestHeapFrame() = %v, want %v", got, tt.want)
			}
		})
	}
}
