// Q1. A) Write a program in GO language to accept two matrices and
// display its multiplication

package main

import (
	"fmt"
)

func main() {
	var rowsA, colsA, rowsB, colsB int

	// Accept dimensions for the first matrix
	fmt.Print("Enter the number of rows for Matrix A: ")
	fmt.Scan(&rowsA)
	fmt.Print("Enter the number of columns for Matrix A: ")
	fmt.Scan(&colsA)

	// Accept dimensions for the second matrix
	fmt.Print("Enter the number of rows for Matrix B: ")
	fmt.Scan(&rowsB)
	fmt.Print("Enter the number of columns for Matrix B: ")
	fmt.Scan(&colsB)

	// Check if matrix multiplication is possible
	if colsA != rowsB {
		fmt.Println("Matrix multiplication is not possible. Number of columns in Matrix A must be equal to the number of rows in Matrix B.")
		return
	}

	// Initialize matrices A and B
	matrixA := make([][]int, rowsA)
	matrixB := make([][]int, rowsB)

	// Accept values for Matrix A
	fmt.Println("\nEnter elements for Matrix A:")
	for i := 0; i < rowsA; i++ {
		matrixA[i] = make([]int, colsA)
		for j := 0; j < colsA; j++ {
			fmt.Printf("Enter element at position (%d, %d): ", i+1, j+1)
			fmt.Scan(&matrixA[i][j])
		}
	}

	// Accept values for Matrix B
	fmt.Println("\nEnter elements for Matrix B:")
	for i := 0; i < rowsB; i++ {
		matrixB[i] = make([]int, colsB)
		for j := 0; j < colsB; j++ {
			fmt.Printf("Enter element at position (%d, %d): ", i+1, j+1)
			fmt.Scan(&matrixB[i][j])
		}
	}

	// Perform matrix multiplication
	result := multiplyMatrices(matrixA, matrixB)

	// Display the result
	fmt.Println("\nResultant Matrix (Matrix A * Matrix B):")
	displayMatrix(result)
}

// Function to perform matrix multiplication
func multiplyMatrices(matrixA, matrixB [][]int) [][]int {
	rowsA, colsA := len(matrixA), len(matrixA[0])
	_, colsB := len(matrixB), len(matrixB[0]) // Fix: Remove rowsB since it is not used

	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	return result
}

// Function to display a matrix
func displayMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println()
	}
}
