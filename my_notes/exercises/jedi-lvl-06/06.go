package main

import "fmt"

func main() {
	func(s string) {
		fmt.Println("hello, ", s)
	}("dude")
}
