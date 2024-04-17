// B) Write a program in the GO language program to open a file in
// READ only mode.
package main

import (
	"fmt"
	"os"
)

func main() {
	// Specify the file path
	filePath := "C:\\Users\\sohai\\OneDrive\\Desktop\\GO_Programming\\slip_no_19\\example.txt"

	// Open the file in READ-only mode
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	fmt.Printf("File '%s' opened successfully in READ-only mode.\n", filePath)

	// You can now read from the file using the 'file' variable.
	// For example, you might use a bufio.Scanner or read the file line by line.
}
