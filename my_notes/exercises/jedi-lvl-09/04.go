package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementor := 0
	gs := 100
	wg.Add(gs)
	var l sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			l.Lock()
			v := incrementor
			v++
			incrementor = v
			fmt.Println(incrementor)
      l.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Done", incrementor)
}
