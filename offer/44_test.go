package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_findNthDigit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 100,
			},
			want: 5,
		},
		{
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			args: args{
				n: 9,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findNthDigit(tt.args.n))
		})
	}
}
