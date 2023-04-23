package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	args: args{
		// 		grid: [][]byte{
		// 			[]byte("11110"),
		// 			[]byte("11010"),
		// 			[]byte("11000"),
		// 			[]byte("00000"),
		// 		},
		// 	},
		// 	want: 1,
		// },
		// {
		// 	args: args{
		// 		grid: [][]byte{
		// 			[]byte("11000"),
		// 			[]byte("11000"),
		// 			[]byte("00100"),
		// 			[]byte("00011"),
		// 		},
		// 	},
		// 	want: 3,
		// },
		{
			args: args{
				grid: [][]byte{
					[]byte("10111"),
					[]byte("10101"),
					[]byte("11101"),
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, numIslands2(tt.args.grid))
		})
	}
}
