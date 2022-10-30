package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go fun("goroutine function call")
	// goroutine with anonymous function
	func() {
		fmt.Println("goroutine with anonymous function")
	}()
	// goroutine with function value call
	go func(value string) {
		fun(value)
	}("goroutine with function value call")
	// wait for goroutines to end

	fmt.Println("done..")
}
