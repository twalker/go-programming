package main

import "fmt"

func doSomething(x int) int {
	return x * 5
}

func main() {
	ch := make(chan int)
  // starting go routine in function literal to gain access to return value of function.
	go func() {
		ch <- doSomething(5) // put return value in channel
	}()
	fmt.Println(<-ch) // pull value out of channel
}
