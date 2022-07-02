package pongo2

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/flosch/pongo2/v5"
	"golang.org/x/exp/constraints"
)

func Test(t *testing.T) {
	tpl, err := pongo2.FromString("Hello {{ name|capfirst }}!")
	if err != nil {
		panic(err)
	}
	out, err := tpl.Execute(pongo2.Context{"name": "florian"})
	if err != nil {
		panic(err)
	}
	fmt.Println(out) // Output: Hello Florian!
}

// 同时还必须在测试函数外定义一个执行泛型用例的泛型函数

func Add[T constraints.Ordered](a, b T) T {
	return a + b
}

// 单元测试函数
func TestAdd(t *testing.T) {
	type testCase[T constraints.Ordered] struct {
		name string
		a    T
		b    T
		want T
	}

	//intTestCases := []testCase[int]{
	//	{
	//		name: "ok",
	//		a:    1,
	//		b:    1,
	//		want: 2,
	//	},
	//	{
	//		name: "ok2",
	//		a:    10,
	//		b:    10,
	//		want: 20,
	//	},
	//}
	//strCases := []testCase[string]{
	//	{
	//		name: "ok",
	//		a:    "A",
	//		b:    "B",
	//		want: "AB",
	//	},
	//	{
	//		name: "ok2",
	//		a:    "Hello",
	//		b:    "World",
	//		want: "HelloWorld",
	//	},
	//}
	//

	type item interface {
		~struct {
			name       string
			a, b, want int
		} |
			~struct {
				name       string
				a, b, want float64
			} |
			~struct {
				name       string
				a, b, want int32
			} |
			~struct {
				name       string
				a, b, want int
			}
	}

	type list[X item] []X

	var myList list[testCase[int]]
	myList = append(myList, testCase[int]{
		name: "ok",
		a:    1,
		b:    1,
		want: 2,
	})

	for _, e := range myList {
		t.Run(e.name, func(t *testing.T) {
			if got := Add(e.a, e.b); !reflect.DeepEqual(got, e.want) {
				t.Errorf("xxxxxxxxxxxxxx")
			}
		})
	}

}

func Test2(t *testing.T) {
	fmt.Println(65537 == 1<<16)

}
