// Q1. A) Write a program in GO language to check whether the accepted
// number is two digit or not.

package main

import (
	"fmt"
)

func main() {
	// Accept a number from the user
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	// Check if the number is a two-digit number
	if isTwoDigit(num) {
		fmt.Printf("%d is a two-digit number.\n", num)
	} else {
		fmt.Printf("%d is not a two-digit number.\n", num)
	}
}

// Function to check if a number is a two-digit number
func isTwoDigit(num int) bool {
	return num >= 10 && num <= 99
}
