// B) Write a program in GO language to read and write Fibonacci
// series to the using channel.

package main

import (
	"fmt"
	"sync"
)

// Function to generate Fibonacci series and write to channel
func generateFibonacci(n int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}

	close(ch)
}

// Function to read and print values from channel
func readChannel(ch <-chan int, done chan<- bool) {
	for num := range ch {
		fmt.Print(num, " ")
	}
	done <- true
}

func main() {
	// Create channels
	fibonacciCh := make(chan int, 10) // buffered channel to store Fibonacci series
	doneCh := make(chan bool)

	// Use WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Start goroutine to generate Fibonacci series
	wg.Add(1)
	go generateFibonacci(10, fibonacciCh, &wg)

	// Start goroutine to read and print values from the channel
	go readChannel(fibonacciCh, doneCh)

	// Wait for both goroutines to finish
	wg.Wait()

	// Close done channel after reading is done
	close(doneCh)

	// Wait for the readChannel goroutine to finish
	<-doneCh
	fmt.Println("\nFibonacci series generated and read successfully.")
}
