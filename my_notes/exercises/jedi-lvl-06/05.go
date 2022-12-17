package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{
		width:  10,
		length: 40,
	}
	info(s)

	c := circle{
		radius: 2,
	}
	info(c)
}
