package main

import (
	"unsafe"
)

const (
	a1 = "abc"
	b1 = 3
	c1 = unsafe.Sizeof(a)
)

const (
	a2 = iota
	b2 = iota
	c2 = iota
)

func main()  {
	println(a2,b2,c2)
	println(a1,b1,c1)
}