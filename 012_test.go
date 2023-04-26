package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, intToRoman(tt.args.num))
		})
	}
}
