package main

import (
	"fmt"
)

func main() {
	type person struct {
		first string
		last  string
		favs  []string
	}

	p1 := person{
		first: "Tim",
		last:  "Walker",
		favs: []string{
			"Rocky road",
			"World class chocolate",
			"Cookies and cream",
		},
	}

	fmt.Println(p1)

	for _, v := range p1.favs {
		fmt.Println(v)
	}
}
