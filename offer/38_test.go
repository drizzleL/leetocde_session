package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_permutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				s: "abc",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, permutation(tt.args.s))
		})
	}
}
