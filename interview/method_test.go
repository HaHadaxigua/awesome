package interview

import (
	"fmt"
	"testing"
)

type a struct {
}

func (*a) Hi() {
	fmt.Println("hi")
}
func (a) Hello() {
	fmt.Println("hello")
}

type ai interface {
	Hi()
	Hello()
}

func TestMethod(t *testing.T) {
	var x ai
	x = new(a)
	x.Hello()
	x.Hi()

	o := a{}
	o.Hello()
	o.Hi()
}
