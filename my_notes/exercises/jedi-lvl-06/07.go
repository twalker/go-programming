package main

import "fmt"

func main() {
	f := func(s string) {
		fmt.Println("hello, ", s)
	}
	f("dude")
}
