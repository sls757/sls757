// B) Write a program in GO language to accept n records of employee
// information (eno,ename,salary) and display records of employees
// having minimum salary.

package main

import (
	"fmt"
)

// Employee structure to store employee information
type Employee struct {
	Eno    int
	Ename  string
	Salary float64
}

func main() {
	var n int
	fmt.Print("Enter the number of records: ")
	fmt.Scan(&n)

	// Create a slice to store n employee records
	employees := make([]Employee, n)

	// Accept employee records
	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for Employee %d:\n", i+1)
		fmt.Print("Employee Number: ")
		fmt.Scan(&employees[i].Eno)
		fmt.Print("Employee Name: ")
		fmt.Scan(&employees[i].Ename)
		fmt.Print("Salary: ")
		fmt.Scan(&employees[i].Salary)
	}

	// Find employee with minimum salary
	minSalary := employees[0].Salary
	for i := 1; i < n; i++ {
		if employees[i].Salary < minSalary {
			minSalary = employees[i].Salary
		}
	}

	// Display employees with minimum salary
	fmt.Println("\nEmployees with Minimum Salary:")
	for i := 0; i < n; i++ {
		if employees[i].Salary == minSalary {
			fmt.Printf("\nEmployee %d:\n", i+1)
			fmt.Printf("Employee Number: %d\n", employees[i].Eno)
			fmt.Printf("Employee Name: %s\n", employees[i].Ename)
			fmt.Printf("Salary: %.2f\n", employees[i].Salary)
		}
	}
}
