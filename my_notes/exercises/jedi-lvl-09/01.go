package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	fmt.Println("Foo")
	wg.Done()
}

func bar() {
	fmt.Println("Bar")
	wg.Done()
}

func main() {
	fmt.Println("start CPU", runtime.NumCPU())
	fmt.Println("start Goroutine", runtime.NumGoroutine())

	wg.Add(2)
	go foo()

	go bar()
	fmt.Println("mid CPU", runtime.NumCPU())
	fmt.Println("mid Goroutine", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("end CPU", runtime.NumCPU())
	fmt.Println("end Goroutine", runtime.NumGoroutine())
}
