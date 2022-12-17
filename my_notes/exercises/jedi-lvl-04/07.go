package main

import "fmt"

func main() {
	// slice of a slice of string
	jb := []string{"James", "Bond"}
	mp := []string{"Money", "Penny"}
	x := [][]string{jb, mp}
	for _, v := range x {
    for _, v2 := range v {
    fmt.Println(v2)
    }
	}
}
