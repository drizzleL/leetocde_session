package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				nums: []int{-2, 0, 1, 1, 2},
			},
			want: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, threeSum(tt.args.nums))
		})
	}
}
