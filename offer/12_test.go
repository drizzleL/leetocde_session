package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				board: [][]byte{
					[]byte("ABCE"),
					[]byte("SFCS"),
					[]byte("ADEE"),
				},
				word: "ABCCED",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, exist(tt.args.board, tt.args.word))
		})
	}
}
