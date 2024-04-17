// Q1. A) Write a program in GO language to print a multiplication table of
// number using function.

package main

import "fmt"

func printMultiplicationTable(number, limit int) {
	fmt.Printf("Multiplication Table for %d:\n", number)
	for i := 1; i <= limit; i++ {
		result := number * i
		fmt.Printf("%d x %d = %d\n", number, i, result)
	}
}

func main() {
	// Specify the number for the multiplication table
	multiplicationNumber := 5

	// Specify the limit (how far the multiplication table should go)
	limit := 10

	// Call the function to print the multiplication table
	printMultiplicationTable(multiplicationNumber, limit)
}
