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

	p2 := person{
		first: "James",
		last:  "Bond",
		favs: []string{
			"Chocolate",
			"Vanilla",
		},
	}

	fmt.Println(p1)

	for _, v := range p1.favs {
		fmt.Println(v)
	}

	m := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}
	for _, p := range m {
		fmt.Println(p.first, p.last, p.favs)
	}
}
