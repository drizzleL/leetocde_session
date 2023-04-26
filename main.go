package main

import (
	"log"
	"math/rand"
)

func main() {
	v := func() func() {
		dict := make([]bool, 10)
		return func() {
			dict[rand.Intn(10)] = true
			log.Println(dict)
		}
	}
	k := v()
	k()
	k()
	k()
	k()
	k()
}
