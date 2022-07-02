package generic

import (
	"fmt"
	"testing"
)

type My interface {
	~string | ~int | ~float64
	HelloMy() string
}

type A string

func (a A) HelloMy() string {
	return "string"
}

type B float64

func (b B) HelloMy() string {
	return "float"
}

type C int

func (c C) HelloMy() string {
	return "int"
}

func Test(t *testing.T) {
	a := new(A)
	b := new(B)
	Do(*a)
	Do(*b)
	Do(*new(C))
}

func Do[T My](x T) {
	fmt.Println(x.HelloMy())
}

func TestAny(t *testing.T) {
	var x any
	x = "12"
	fmt.Println(x)
}
