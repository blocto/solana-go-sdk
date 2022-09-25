package pointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want *uint64
	}{
		{
			args: args{
				v: 1,
			},
			want: Get[uint64](1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Get[uint64](1))
		})
	}
}
