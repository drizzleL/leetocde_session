package main

import (
	"log"
	"testing"
)

func Test41(t *testing.T) {
	m := ConstructorM()
	m.AddNum(-1)
	log.Println(m.FindMedian())
	m.AddNum(-2)
	log.Println(m.FindMedian())
	m.AddNum(-3)
	log.Println(m.FindMedian())
	m.AddNum(-4)
	log.Println(m.FindMedian())
	m.AddNum(-5)
	log.Println(m.FindMedian())
}
