// Q1. A) Write a program in GO language program to create Text file 

package main

import (
	"fmt"
	"os"
)

func main() {
	// Open or create a text file for writing
	file, err := os.Create("example.txt")

	// Check for errors when creating the file
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // defer closing the file until the surrounding function returns

	// Write content to the file
	content := "Hello, this is a sample text file created using Go!"
	_, err = file.WriteString(content)

	// Check for errors when writing to the file
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Text file created successfully.")
}
