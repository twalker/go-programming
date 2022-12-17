package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s := even(sum, ii...)
	s2 := odd(sum, ii...)
	fmt.Println("even", s)
	fmt.Println("odd", s2)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

/*cb func signature */ /*other parameters*/ /*return type*/
func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(cb func(xi ...int) int, vi ...int) int {
	var oi []int
	for _, v := range vi {
		if v%2 == 1 {
			oi = append(oi, v)
		}
	}
	return cb(oi...)
}
