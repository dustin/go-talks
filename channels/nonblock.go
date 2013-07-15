package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	select {
	case n := <-ch:
		fmt.Printf("I read %v from ch\n", n)
	default:
		fmt.Printf("There's no data on ch\n")
	}
}
