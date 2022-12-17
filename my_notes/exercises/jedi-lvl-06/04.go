package main

import "fmt"

func main() {
	p := person{
		first: "Tim",
		last:  "Walker",
		age:   53,
	}
  p.speak()
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("%v %v is %d years old\n", p.first, p.last, p.age)
}
