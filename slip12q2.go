// B) Write a program in GO language that creates a slice of integers,
// checks numbers from the slice are even or odd and further sent to
// respective go routines through channel and display values
// received by goroutines.

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a slice of integers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Create channels for even and odd numbers
	evenCh := make(chan int)
	oddCh := make(chan int)

	// Use WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Start two goroutines to process even and odd numbers
	wg.Add(2)
	go processNumbers(numbers, evenCh, oddCh, &wg)

	// Start a goroutine to display even numbers
	go displayNumbers(evenCh, "Even", &wg)

	// Start a goroutine to display odd numbers
	go displayNumbers(oddCh, "Odd", &wg)

	// Wait for all goroutines to finish
	wg.Wait()
}

// Function to process numbers and send them to respective channels
func processNumbers(numbers []int, evenCh chan<- int, oddCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, num := range numbers {
		if num%2 == 0 {
			evenCh <- num
		} else {
			oddCh <- num
		}
	}

	// Close channels for even and odd numbers after processing is done
	close(evenCh)
	close(oddCh)
}

// Function to display numbers received from the channel
func displayNumbers(ch <-chan int, label string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s Numbers: ", label)
	for num := range ch {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
