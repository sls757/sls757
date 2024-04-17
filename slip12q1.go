// Q1. A) Write a program in GO language to swap two numbers using call
// by reference concept

package main

import "fmt"

// Function to swap two numbers using call by reference
func swapByReference(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	// Declare and initialize two numbers
	num1 := 5
	num2 := 10

	// Display the original values
	fmt.Printf("Before swapping: num1 = %d, num2 = %d\n", num1, num2)

	// Call the swapByReference function to swap the numbers
	swapByReference(&num1, &num2)

	// Display the swapped values
	fmt.Printf("After swapping: num1 = %d, num2 = %d\n", num1, num2)
}
