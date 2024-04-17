// B) Write a program in GO language to copy all elements of one array
// into another using a method.

package main

import (
	"fmt"
)

// CopyArray method to copy elements of one array into another
func CopyArray(source []int, destination []int) {
	for i, value := range source {
		destination[i] = value
	}
}

func main() {
	var n int

	fmt.Print("Enter the size of the array: ")
	fmt.Scan(&n)

	// Initialize source array
	sourceArray := make([]int, n)

	// Accept elements for the source array
	fmt.Println("Enter elements for the source array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Element %d: ", i+1)
		fmt.Scan(&sourceArray[i])
	}

	// Initialize destination array with the same size as the source array
	destinationArray := make([]int, n)

	// Call the CopyArray method to copy elements from source to destination
	CopyArray(sourceArray, destinationArray)

	// Display the elements of the destination array
	fmt.Println("\nElements of the Destination Array (copied from Source Array):")
	for _, value := range destinationArray {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}
