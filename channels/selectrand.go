package main

import (
	"fmt"
)

func generate(n int) <-chan int {
	rv := make(chan int)
	go func() {
		for {
			rv <- n
		}
	}()
	return rv
}

func main() {
	ch1 := generate(7)
	ch2 := generate(11)
	for i := 0; i < 8; i++ {
		n := 0
		select {
		case n = <-ch1:
		case n = <-ch2:
		}
		fmt.Printf("Iteration %v got %v\n", i, n)
	}
}
