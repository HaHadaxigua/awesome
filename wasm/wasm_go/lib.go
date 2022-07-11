package main

import "fmt"

func main() {}

//export Hello
func Hello() {
	fmt.Println("hello world")
}
