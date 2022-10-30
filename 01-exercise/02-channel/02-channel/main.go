package main

import (
	"fmt"
)

//var wg sync.WaitGroup

func main() {
	c := make(chan int)
	//wg.Add(1)
	go func() {
		//defer wg.Done()
		defer close(c)
		for i := 0; i < 6; i++ {
			// TODO: send iterator over channel
			c <- i
		}
	}()
	//wg.Wait()
	// TODO: range over channel to recv values
	for val := range c {
		fmt.Println(val)
	}

}
