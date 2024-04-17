// B) Write a program in GO language to sort array elements in
// ascending order.

package main

import (
	"fmt"
)

// Function to perform bubble sort on an array
func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap elements if they are in the wrong order
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	var n int
	fmt.Print("Enter the size of the array: ")
	fmt.Scan(&n)

	// Initialize an array of size n
	arr := make([]int, n)

	// Accept array elements
	fmt.Println("Enter the array elements:")
	for i := 0; i < n; i++ {
		fmt.Printf("Element %d: ", i+1)
		fmt.Scan(&arr[i])
	}

	// Call the bubbleSort function to sort the array
	bubbleSort(arr)

	// Display the sorted array
	fmt.Println("\nSorted Array (Ascending Order):")
	for _, element := range arr {
		fmt.Printf("%d ", element)
	}
	fmt.Println()
}
