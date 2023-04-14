package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_isStraight(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				[]int{0, 0, 1, 2, 5},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isStraight(tt.args.nums))
		})
	}
}
