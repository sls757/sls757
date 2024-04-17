// Q1. A) Write a program in GO language to create an interface and display
// its values with the help of type assertion.

package main

import "fmt"

// Define an interface named MyInterface
type MyInterface interface {
	display()
}

// Define a struct type implementing MyInterface
type MyStruct struct {
	Value string
}

// Implement the display method for MyStruct
func (ms MyStruct) display() {
	fmt.Printf("Value from MyStruct: %s\n", ms.Value)
}

// Define another struct type implementing MyInterface
type AnotherStruct struct {
	Number int
}

// Implement the display method for AnotherStruct
func (as AnotherStruct) display() {
	fmt.Printf("Number from AnotherStruct: %d\n", as.Number)
}

func main() {
	// Create instances of MyStruct and AnotherStruct
	myInstance := MyStruct{Value: "Hello, World!"}
	anotherInstance := AnotherStruct{Number: 42}

	// Declare variables of type MyInterface
	var interfaceVar MyInterface

	// Assign MyStruct instance to MyInterface variable
	interfaceVar = myInstance
	displayWithTypeAssertion(interfaceVar)

	// Assign AnotherStruct instance to MyInterface variable
	interfaceVar = anotherInstance
	displayWithTypeAssertion(interfaceVar)
}

// Function to display values with type assertion
func displayWithTypeAssertion(i MyInterface) {
	// Use type assertion to determine the underlying type
	switch v := i.(type) {
	case MyStruct:
		v.display()
	case AnotherStruct:
		v.display()
	default:
		fmt.Println("Unknown type")
	}
}
