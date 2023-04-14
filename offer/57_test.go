package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums:   []int{14, 15, 16, 22, 53, 60},
				target: 76,
			},
			want: []int{2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, twoSum(tt.args.nums, tt.args.target))
		})
	}
}
