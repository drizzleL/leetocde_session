package main

import (
	"log"
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {
	input := []int{2, 3, 1, 0, 2, 5, 3}
	log.Println(findRepeatNumber(input))

}
