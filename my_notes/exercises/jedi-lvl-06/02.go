package main

import "fmt"

func main() {
	nums := []int{2, 4, 6}
	n := foo(nums...)
	m := bar(nums)
	fmt.Println(n, m)
}

func foo(n ...int) int {
	tot := 0
	for _, v := range n {
		tot += v
	}
	return tot
}

func bar(n []int) int {
	tot := 0
	for _, v := range n {
		tot += v
	}
	return tot
}
