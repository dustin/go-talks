package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	select {
	case n := <-ch:
		fmt.Printf("I read %v from ch\n", n)
	case <-time.After(time.Millisecond):
		fmt.Printf("Timed out after 1ms\n")
	}
}
