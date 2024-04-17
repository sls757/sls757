// Q1. A) Write a program in GO language to print a recursive sum of digits
// of a given number

package main

import (
	"fmt"
)

// Function to calculate the sum of digits recursively
func recursiveSumOfDigits(num int) int {
	if num == 0 {
		return 0
	}
	return num%10 + recursiveSumOfDigits(num/10)
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Calculate the recursive sum of digits
	sum := recursiveSumOfDigits(num)

	// Display the result
	fmt.Printf("Recursive Sum of Digits for %d: %d\n", num, sum)
}
