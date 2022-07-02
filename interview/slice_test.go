package interview

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	a := []int{
		4, 5, 6, 7,
	}
	wannerAdd := []int{
		1, 2, 3,
	}
	//a = append(wannerAdd, a...)

	for i := range wannerAdd {
		if i < len(a) {
			a[i], wannerAdd[i] = wannerAdd[i], a[i]
		}
	}
	a = append(a[:len(wannerAdd)], wannerAdd...)

}

func add(a []int) {
	a = append(a, 4)
}

func TestCompare(t *testing.T) {
	//type A Base
	//a := &A{}
	//fmt.Println(a.Hello())
	DeferFunc2 := func(i int) int {
		x := i
		defer func() {
			x = 3 + x
		}()
		return x
	}

	DeferFunc1 := func(i int) (x int) {
		defer func() {
			x += 3
		}()
		return x
	}
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
}

type Base struct {
}

func (b *Base) Hello() string {
	return "hello"
}
