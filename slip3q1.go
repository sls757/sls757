// Q1. A)
// Max Marks: 35+15=50
// Write a program in the GO language using function to check
// whether accepts number is palindrome or not.

package main

import (
	"fmt"
)

func main() {
	var no int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&no)
	if isPalindrome(no) {
		fmt.Print(no, " is a palindrome number")
	} else {
		fmt.Print(no, " is not a palindrome number")
	}
}

func isPalindrome(no int) bool {
	originalNumber := no
	reversedNumber := 0

	for no > 0 {
		remainder := no % 10
		reversedNumber = reversedNumber*10 + remainder
		no /= 10
	}

	return originalNumber == reversedNumber
}
