package main

import "fmt"

// TODO: Implement relaying of message with Channel Direction

func genMsg(ch1 chan string, message ...string) {
	for _, val := range message {
		ch1 <- val
	}
	close(ch1)
	// send message on ch1
}

func relayMsg(ch1 chan string, ch2 chan string) {
	for val := range ch1 {
		ch2 <- val
	}

	close(ch2)
	// recv message on ch1
	// send it on ch2
}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// spine goroutine genMsg and relayMsg
	go genMsg(ch1, "hello", "world", "!")
	go relayMsg(ch1, ch2)

	// recv message on ch2
	for val := range ch2 {
		fmt.Println(val)
	}
}
