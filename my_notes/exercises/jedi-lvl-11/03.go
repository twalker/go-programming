package main

import "fmt"

type customErr struct{
  info string
}

func (ce customErr) Error() string {
  return fmt.Sprintf("My custom error: %v", ce.info)
}

func foo(e error) {
  fmt.Println(e)
  // assertion into custom error --different from conversion
  fmt.Println(e.(customErr).info)
}

func main() {
  c1 := customErr{ 
    info: "my custom error data",
  }
  foo(c1)
}
