// B) Write a program in GO language to add or append content at the
// end of a text file//not understand

package main

import (
	"fmt"
	"os"
	"log"
)

func appendToFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	fmt.Printf("Content appended to %s\n", filename)
	return nil
}

func main() {
	// Specify the file path and content to append
	filePath := "example.txt"
	contentToAppend := "This content will be appended to the file.\n"

	// Call the function to append content to the file
	err := appendToFile(filePath, contentToAppend)
	if err != nil {
		log.Fatal(err)
	}
}
