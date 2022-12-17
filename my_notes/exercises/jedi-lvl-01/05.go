package main

import (
	"fmt"
)

type hotdog int

var y int

func main() {
	var x hotdog
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
