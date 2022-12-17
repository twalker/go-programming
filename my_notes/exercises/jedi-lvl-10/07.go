package main

import "fmt"

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		// go goForIt(c)
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}
  for k := 0; k < 100; k++ {
    fmt.Println(k, <-c)
  }
}

// func goForIt(c <-chan ){
//   for i:= 0; i < 10; i++ {
//     c <- i
//   }
// }
