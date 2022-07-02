package script

import (
	"fmt"

	"awesome/yaegi/linux_complex/types"
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

func useOutsideStruct() {
	a := types.Os{
		Cloud:    "aws",
		Platform: "ubuntu",
	}
	fmt.Println("use outside struct: ", a)
	fmt.Println("use outside value: ", types.AwsOs)
}

type Result struct {
	Code    int
	Message string
}

func GiveInvoker() *Result {
	return &Result{
		Code:    11111,
		Message: "let invoker check",
	}
}

func Main() {
	fmt.Println(Fib(10))
	slice := cast.ToStringSlice(Fib(10))
	fmt.Println("slice: ", slice)

	useOutsideStruct()
}
