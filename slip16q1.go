// Q1. A) Write a program in GO language to create a user defined package
// to find out the area of a rectangle.

// main.go

package main

import (
	"fmt"
	"/home/code/Desktop/College_Practical_2023/GO_Programming/slip_no_16/geometry/rectangle.go" // Replace with the actual path to your geometry package
)

func main() {
	// Example usage of the user-defined package
	length := 5.0
	width := 3.0

	// Calculate the area using the AreaOfRectangle function from the geometry package
	area := geometry.AreaOfRectangle(length, width)

	// Print the result
	fmt.Printf("Area of rectangle with length %.2f and width %.2f: %.2f\n", length, width, area)
}
