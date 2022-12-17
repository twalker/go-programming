package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)

	fmt.Printf("%T", x)
	fmt.Println(x)
	fmt.Printf("%T", y)
	fmt.Println(y)
	fmt.Printf("%T", z)
	fmt.Println(z)
}
