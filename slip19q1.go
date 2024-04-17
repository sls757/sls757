// Q1. A) Write a program in GO language to illustrate the function
// returning multiple values(add, subtract).

package main

import "fmt"

func performAddAndSubtract(a, b int) (int, int) {
	// Addition
	sum := a + b

	// Subtraction
	diff := a - b

	return sum, diff
}

func main() {
	// Example values
	num1 := 10
	num2 := 5

	// Calling the function and receiving multiple values
	additionResult, subtractionResult := performAddAndSubtract(num1, num2)

	// Displaying the results
	fmt.Printf("Addition: %d\n", additionResult)
	fmt.Printf("Subtraction: %d\n", subtractionResult)
}
