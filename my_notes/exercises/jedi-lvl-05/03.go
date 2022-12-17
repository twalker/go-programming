package main

import "fmt"

func main() {
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}
	type sedan struct {
		vehicle
		luxury bool
	}

	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Black",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(s)
	// inner type promotion
	fmt.Println(t.color, t.fourWheel)
	fmt.Println(s.color, s.luxury)
}
