package main

import (
	"fmt"
	"runtime"
)

func main() {
  fmt.Println("Arch", runtime.GOARCH)
  fmt.Println("Arch", runtime.GOOS)
}

