package main

import "fmt"

func main() {
	fmt.Println("Hello")
	foo(func() int {
		return 74
	}, "Hi")
}

func foo(cb func() int, s string) {
	fmt.Println("recieved", s, "about to call callback")
	cb()
}
