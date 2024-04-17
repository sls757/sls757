// B) Write a program in Go language how to create a channel and
// illustrate how to close a channel using for range loop and close
// function.

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a channel of integers
	intChannel := make(chan int)

	// Use a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Producer: Generate and send numbers to the channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			intChannel <- i
		}
		// Close the channel when done producing
		close(intChannel)
	}()

	// Consumer: Receive numbers from the channel using for range loop
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range intChannel {
			fmt.Printf("Received: %d\n", num)
		}
	}()

	// Wait for both goroutines to finish
	wg.Wait()
}
