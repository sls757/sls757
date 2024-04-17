// B) Write a program in GO language to read XML file into structure
// and display structure

package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Define a structure to represent the XML data
type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	City    string   `xml:"city"`
}

func main() {
	// Open and read the XML file
	file, err := os.Open("person.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a variable of the structure type to hold the data
	var person Person

	// Use the XML decoder to decode the XML file into the structure
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return
	}

	// Display the structure
	fmt.Println("Person Information:")
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("City: %s\n", person.City)
}
