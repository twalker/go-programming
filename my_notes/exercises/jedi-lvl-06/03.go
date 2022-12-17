package main

import "fmt"

func main() {
	fmt.Println("starting")
	defer foo()
	fmt.Println("ending")
}

func foo() {
	defer func() {
		fmt.Println("foo defer ran")
	}()
	fmt.Println("foo ran")
}
