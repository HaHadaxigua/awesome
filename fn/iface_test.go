package main

import "testing"

type (
	a int
	b int
)

func (a) String() {
	println("a")
}

func (b) String() {
	println("b")
}

func Test(t *testing.T) {
	method := interface{ String() }.String
	var aa a
	method(aa)
	method(b(1))
}
