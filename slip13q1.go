// Q1. A) Write a program in GO language to print sum of all even and odd
// numbers separately between 1 to 100.

package main

import "fmt"

func main() {
	// Initialize sum variables for even and odd numbers
	evenSum := 0
	oddSum := 0

	// Calculate sum of even and odd numbers between 1 and 100
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			evenSum += i
		} else {
			oddSum += i
		}
	}

	// Print the sum of even and odd numbers
	fmt.Printf("Sum of even numbers between 1 and 100: %d\n", evenSum)
	fmt.Printf("Sum of odd numbers between 1 and 100: %d\n", oddSum)
}
