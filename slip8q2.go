// B) Write a program in GO language to create an interface shape that
// includes area and perimeter. Implements these methods in circle
// and rectangle type.

package main

import (
	"fmt"
	"math"
)

// Define the Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define the Circle type
type Circle struct {
	Radius float64
}

// Implement the Shape interface for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
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

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func main() {
	// Create a Circle
	circle := Circle{Radius: 5}

	// Create a Rectangle
	rectangle := Rectangle{Width: 4, Height: 6}

	// Use the Shape interface to calculate area and perimeter
	printShapeDetails(circle)
	printShapeDetails(rectangle)
}

// Function to print details of a shape using the Shape interface
func printShapeDetails(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
	fmt.Println()
}
