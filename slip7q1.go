// Q1. A) Write a program in GO language to accept one matrix and display
// its transpose.

package main

import (
	"fmt"
)

func main() {
	// Input matrix
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Display original matrix
	fmt.Println("Original Matrix:")
	displayMatrix(matrix)

	// Get and display transpose of the matrix
	transpose := getTranspose(matrix)
	fmt.Println("\nTranspose Matrix:")
	displayMatrix(transpose)
}

// Function to display a matrix
func displayMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, element := range row {
			fmt.Printf("%d\t", element)
		}
		fmt.Println()
	}
}

// Function to calculate the transpose of a matrix
func getTranspose(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	// Create a new matrix for the transpose
	transpose := make([][]int, cols)
	for i := range transpose {
		transpose[i] = make([]int, rows)
	}

	// Calculate the transpose
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transpose[j][i] = matrix[i][j]
		}
	}

	return transpose
}
