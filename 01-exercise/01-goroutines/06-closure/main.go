package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// what is the output
	//TODO: fix the issue.

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Println(val)
		}(i)
	}
	wg.Wait()
}
