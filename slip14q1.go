// Q1. A) Write a program in GO language to demonstrate working of slices
// (like append, remove, copy etc.)

package main

import (
	"fmt"
)

func main() {
	// Create a slice with some initial elements
	mySlice := []int{1, 2, 3, 4, 5}

	// Append elements to the slice
	mySlice = append(mySlice, 6, 7, 8)
	fmt.Println("After append:", mySlice)

	// Remove an element from the slice by creating a new slice excluding the element
	indexToRemove := 2
	if indexToRemove >= 0 && indexToRemove < len(mySlice) {
		mySlice = append(mySlice[:indexToRemove], mySlice[indexToRemove+1:]...)
		fmt.Println("After remove:", mySlice)
	} else {
		fmt.Println("Invalid index to remove.")
	}

	// Copy the slice to a new slice
	copyOfSlice := make([]int, len(mySlice))
	copy(copyOfSlice, mySlice)
	fmt.Println("Copy of the slice:", copyOfSlice)
}
