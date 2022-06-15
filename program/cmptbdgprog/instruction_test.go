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

func TestSetComputeUnitLimit(t *testing.T) {
	type args struct {
		param SetComputeUnitLimitParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: SetComputeUnitLimitParam{
					Units: 1000,
				},
			},
			want: types.Instruction{
				ProgramID: common.ComputeBudgetProgramID,
				Accounts:  []types.AccountMeta{},
				Data:      []byte{2, 232, 3, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetComputeUnitLimit(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetComputeUnitLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetComputeUnitPrice(t *testing.T) {
	type args struct {
		param SetComputeUnitPriceParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: SetComputeUnitPriceParam{
					MicroLamports: 1000,
				},
			},
			want: types.Instruction{
				ProgramID: common.ComputeBudgetProgramID,
				Accounts:  []types.AccountMeta{},
				Data:      []byte{3, 232, 3, 0, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetComputeUnitPrice(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetComputeUnitPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
