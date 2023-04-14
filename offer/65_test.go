package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				a: 100,
				b: 22,
			},
			want: 122,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, add(tt.args.a, tt.args.b))
		})
	}
}
