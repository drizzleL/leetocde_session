package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dicesProbability(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			args: args{
				2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, dicesProbability(tt.args.n))
		})
	}
}
