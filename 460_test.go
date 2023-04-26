package main

import "testing"

func TestLfuCache(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Get(1)
	c.Put(3, 3)
	c.Get(2)
	c.Get(3)
	c.Put(4, 4)
	c.Get(1)
	c.Get(3)
	c.Get(4)
}
func TestLfuCache2(t *testing.T) {
	c := Constructor(2)
	c.Put(3, 1)
	c.Put(2, 1)
	c.Put(2, 2)
	c.Put(4, 4)
	c.Get(2)
}
func TestLfuCache3(t *testing.T) {
	c := Constructor(1)
	c.Put(2, 1)
	c.Get(2)
	c.Put(3, 2)
	c.Get(2)
	c.Get(3)
}
