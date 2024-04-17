// Q1. A) Write a program in GO language to demonstrate function return
// multiple values.

package main

import "fmt"

// Function that returns the sum and difference of two numbers
func sumAndDifference(a, b int) (int, int) {
	sum := a + b
	difference := a - b
	return sum, difference
}

func main() {
	// Call the function and receive multiple return values
	resultSum, resultDiff := sumAndDifference(10, 5)

	// Print the results
	fmt.Printf("Sum: %d\n", resultSum)
	fmt.Printf("Difference: %d\n", resultDiff)
}
