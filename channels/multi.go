package main

import "fmt"

const nworkers = 8

func summer(chin, chout chan int) {
	total := 0
	for n := range chin {
		total += n
	}
	chout <- total
}

func main() {
	workch := make(chan int)
	resch := make(chan int)

	for i := 0; i < nworkers; i++ {
		go summer(workch, resch) // HL
	}

	for i := 0; i < 100; i++ {
		workch <- i // HL
	}
	close(workch)

	total := 0
	for i := 0; i < nworkers; i++ {
		n := <-resch // HL
		fmt.Printf("Worker finished with %v\n", n)
		total += n
	}
	fmt.Printf("Total is %v\n", total)
}
