package main

import (
	"fmt"
)

func main() {
	foo := struct {
		first string
		cars  map[string]int
		favs  []string
	}{
		first: "Tim",
		cars: map[string]int{
			"honda": 2,
			"vw":    1,
		},
		favs: []string{"vanilla", "choc"},
	}

	fmt.Println(foo)
	fmt.Println(foo.cars)
	for k, v := range foo.favs {
		fmt.Println(k, v)
	}
}
