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

func processInts(in <-chan int, wg *sync.WaitGroup) {
	go func() {
		fmt.Println("Process starting...")
		for n := range in {
			fmt.Println("-> Got", n)
		}
		fmt.Println("Process done.")
		defer wg.Done()
	}()
}

func main() {
	var wg sync.WaitGroup

	// Set up the pipeline.
	c := generateInts(20)
	processInts(c, &wg)
	wg.Add(1)

	fmt.Println("***** WAIT *****")
	wg.Wait()
	fmt.Println("***** DONE *****")
}
