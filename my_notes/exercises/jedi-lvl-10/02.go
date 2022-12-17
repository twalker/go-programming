package main

import (
	"fmt"
)

func main() {
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
	// cs := make(chan int)

	// go func() {
	// 	cs <- 42
	// }()
	// fmt.Println(<-cs)

	// fmt.Printf("------\n")
	// fmt.Printf("cs\t%T\n", cs)
}
