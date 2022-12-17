package main

import "fmt"

func foo() func() int {
	return func() int {
		fmt.Println("called")
		return 42
	}
}

func main() {
	f := foo()
	fmt.Println(f())
}
