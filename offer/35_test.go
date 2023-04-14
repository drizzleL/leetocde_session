package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				[]int{1, 3, 5, 6},
				2,
			},
			want: 1,
		},
		{
			args: args{
				[]int{1, 3, 5, 6},
				7,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, searchInsert(tt.args.nums, tt.args.target))
		})
	}
}
