package common

import (
	"reflect"
	"testing"
)

func TestUintToVarLenBytes(t *testing.T) {
	type args struct {
		l uint64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{l: 127},
			want: []byte{0x7f},
		},
		{
			args: args{l: 128},
			want: []byte{0x80, 0x01},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintToVarLenBytes(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64ToVarLenBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
