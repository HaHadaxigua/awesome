package main

import (
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"reflect"
)

//go:generate go install github.com/traefik/yaegi/cmd/yaegi@latest
//go:generate yaegi extract github.com/spf13/cast

var Symbols = map[string]map[string]reflect.Value{}

func main() {
	intp := interp.New(interp.Options{}) // 初始化一个 yaegi 解释器
	err := intp.Use(stdlib.Symbols)      // 允许脚本调用（几乎）所有的 Go 官方 package 代码
	if err != nil {
		panic(err)
	}
	intp.Use(Symbols)

	src := `
package script

import (
  "fmt"
  "github.com/spf13/cast"
)

func Fib(n int) int {
	return fib(n, 0, 1)
}

func fib(n, a, b int) int {
	if n == 0 {
		return a
	} else if n == 1 {
		return b
	}
	return fib(n-1, b, a+b)
}

func Main(){
	Fib(30)
}

func main() {
	v := Fib(30)
	fmt.Println(v)
	slice := cast.ToStringSlice(v)
	fmt.Println(slice)
}
`

	_, err = intp.Eval(src)
	if err != nil {
		panic(err)
	}

	v, err := intp.Eval("script.main")
	if err != nil {
		panic(err)
	}

	fu := v.Interface().(func())
	fu()
}
