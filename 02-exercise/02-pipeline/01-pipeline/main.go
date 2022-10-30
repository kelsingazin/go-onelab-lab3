package main

import "fmt"

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	intsCh := make(chan int)

	go func() {
		defer close(intsCh)
		for _, val := range nums {
			intsCh <- val
		}
	}()

	return intsCh
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		for val := range in {
			output <- val * val
		}
	}()

	return output
}

func main() {
	// set up the pipeline
	in := generator(1, 2, 3, 4, 5)

	// run the last stage of pipeline
	out := square(in)
	// receive the values from square stage
	// print each one, until channel is closed.
	for val := range out {
		fmt.Println(val)
	}

}
