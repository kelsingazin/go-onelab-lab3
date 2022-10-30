package main

import "fmt"

func main() {

	intCh := make(chan int)

	go func(a, b int) {
		c := a + b
		intCh <- c
	}(1, 2)
	// TODO: get the value computed from goroutine
	fmt.Printf("computed value %v\n", <-intCh)
}
