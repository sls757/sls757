// B)
// Write a program in GO language to print file information.

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Provide the path to the file for which you want to print information
	filePath := "slips.pdf"

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File Information:")
	fmt.Println("Name:", fileInfo.Name())
	fmt.Println("Size (in bytes):", fileInfo.Size())
	fmt.Println("Mode:", fileInfo.Mode())
	fmt.Println("Last Modified:", fileInfo.ModTime().Format(time.RFC3339))
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Println("System ID (device number):", fileInfo.Sys())
}
