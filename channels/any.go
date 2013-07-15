package main

import (
	"fmt"
)

func worker(name string, ch chan string) {
	for _ = range ch {
		fmt.Printf("Message picked up by %v\n", name)
	}
}

func main() {
	ch1 := make(chan string)
	go worker("worker 1", ch1)
	ch2 := make(chan string)
	go worker("worker 2", ch2)
	ch3 := make(chan string)
	go worker("worker 3", ch3)

	for i := 0; i < 8; i++ {
		sendToAny(ch1, ch2, ch3, "hi")
	}
}

func sendToAny(ch1, ch2, ch3 chan string, task string) {
	select {
	case ch1 <- task:
	case ch2 <- task:
	case ch3 <- task:
	}
}
