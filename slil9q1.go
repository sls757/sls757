// Q1. A) Write a program in GO language using a function to check
// whether the accepted number is palindrome or not

package main

import (
	"fmt"
	"strconv"
)

// Function to check if a number is a palindrome
func isPalindrome(num int) bool {
	// Convert the number to a string
	numStr := strconv.Itoa(num)

	// Compare the string with its reverse
	for i, j := 0, len(numStr)-1; i < j; i, j = i+1, j-1 {
		if numStr[i] != numStr[j] {
			return false
		}
	}

	return true
}

func main() {
	// Accept a number from the user
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	// Check if the number is a palindrome
	if isPalindrome(num) {
		fmt.Printf("%d is a palindrome.\n", num)
	} else {
		fmt.Printf("%d is not a palindrome.\n", num)
	}
}
