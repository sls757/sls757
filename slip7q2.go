// B) Write a program in GO language to create structure student. Writea
// method show() whose receiver is a pointer of struct student

package main

import "fmt"

// Define a Student structure
type Student struct {
	ID    int
	Name  string
	Grade float64
}

// Method to display information about the student
func (s *Student) show() {
	fmt.Printf("Student ID: %d\n", s.ID)
	fmt.Printf("Student Name: %s\n", s.Name)
	fmt.Printf("Student Grade: %.2f\n", s.Grade)
}

func main() {
	// Create a new student using the structure
	student1 := Student{
		ID:    1,
		Name:  "John Doe",
		Grade: 85.5,
	}

	// Call the show method to display information about the student
	fmt.Println("Information about the student:")
	student1.show()
}
