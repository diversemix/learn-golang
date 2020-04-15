package main

import (
	"fmt"
	"sync"
	"time"
)

func generateInts(max int) <-chan int {
	out := make(chan int)
	go func() {
		fmt.Println("Generation starting...")
		for i := 0; i < max; i++ {
			time.Sleep(10 * time.Millisecond)
			fmt.Print("Generated ", i, "->")
			out <- i
		}
		fmt.Println("Generation done.")

		close(out)
	}()
	return out
}

func fanoutInts(in <-chan int, OutA, OutB chan int) {
	for data := range in {
		select {
		case OutA <- data:
		case OutB <- data:
		}
		fmt.Println("Fanout", data)
	}
	close(OutA)
	close(OutB)
}

func processInts(name string, in <-chan int, wg *sync.WaitGroup) {
	go func() {
		fmt.Println(name, "Process starting...")
		for n := range in {
			fmt.Println(name, "-> Got", n)
		}
		fmt.Println(name, "Process done.")
		defer wg.Done()
	}()
}

func main() {
	var wg sync.WaitGroup

	// Set up the pipeline.
	c := generateInts(20)
	outA := make(chan int)
	outB := make(chan int)
	go fanoutInts(c, outA, outB)

	processInts("A", outA, &wg)
	processInts("B", outB, &wg)
	wg.Add(2)

	fmt.Println("***** WAIT *****")
	wg.Wait()
	fmt.Println("***** DONE *****")
}
