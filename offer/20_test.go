package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "0",
			},
			want: true,
		},
		{
			args: args{
				s: "e",
			},
			want: false,
		},
		{
			args: args{
				s: ".",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isNumber(tt.args.s))
		})
	}
}
