package main

import "fmt"

func main() {
	var foo int = 12
	fmt.Printf("%d %b %#x\n", foo, foo, foo)
	var bar = foo << 1
	fmt.Printf("%d %b %#x\n", bar, bar, bar)
}
