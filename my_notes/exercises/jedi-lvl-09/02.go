package main

import "fmt"

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hi, My name is ", p.name)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	fmt.Println("Hello")
	p1 := person{
		name: "Tim",
	}
	// p1.speak()
	// saySomething(p1)
	saySomething(&p1)
}
