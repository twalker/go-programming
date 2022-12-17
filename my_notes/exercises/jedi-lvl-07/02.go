package main

import "fmt"

func main() {
	p := person{
		first: "Tim",
	}
	fmt.Println(p)
	p.changeMe("Jack")
	fmt.Println(p)
	
  changeMe(&p, "Fred")
	fmt.Println(p)
}

type person struct {
	first string
}
// change in receiver function
func (p *person) changeMe(n string) {
	(*p).first = n
}

// change in regular function
func changeMe(p *person, n string) {
	(*p).first = n  // OR just p.first = n 
}
