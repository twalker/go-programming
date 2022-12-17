package main

import "fmt"

func main() {
	m := map[string][]string{
		"James":      {"foo", "bar"},
		"Moneypenny": {"baz", "qux"},
	}

	fmt.Println(m)

	m["Tim"] = []string{"blah", "blast", "off"}
	delete(m, "James")

	for k, v := range m {
		fmt.Printf("%v\n", k)
		for _, val := range v {
			fmt.Printf("\t\t%v\n", val)
		}
	}
}
