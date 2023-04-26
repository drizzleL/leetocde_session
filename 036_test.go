package main

import "testing"

func TestSudo(t *testing.T) {
	v := [][]byte{
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("4.......4"),
	}
	isValidSudoku(v)
}
