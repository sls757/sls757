// B) Write a program in GO language to create an interface shape that
// includes area and volume. Implements these methods in square
// and rectangle type.

package main

import (
	"fmt"
)

// Define the Shape interface
type Shape interface {
	Area() float64
	Volume() float64
}

// Define the Square type
type Square struct {
	SideLength float64
}

// Implement the Shape interface for Square
func (s Square) Area() float64 {
	return s.SideLength * s.SideLength
}

func (s Square) Volume() float64 {
	// Volume is not applicable for 2D shapes, return 0
	return 0
}

// Define the Rectangle type
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Volume() float64 {
	// Volume is not applicable for 2D shapes, return 0
	return 0
}

func main() {
	// Create a Square
	square := Square{SideLength: 4}

	// Create a Rectangle
	rectangle := Rectangle{Width: 3, Height: 5}

	// Use the Shape interface to calculate area and volume
	printShapeDetails(square)
	printShapeDetails(rectangle)
}

// Function to print details of a shape using the Shape interface
func printShapeDetails(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Volume: %.2f\n", s.Volume())
	fmt.Println()
}
