package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Last  string
	Age   int
}

type ByName []person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}
	people := []person{p1, p2}
	fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println(people)
}
