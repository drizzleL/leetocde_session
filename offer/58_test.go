package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "the lazy dog",
			},
			want: "dog lazy the",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseWords(tt.args.s))
		})
	}
}
