// B) Write a Program in GO language to accept n records of employee [20 Marks]
// information (eno,ename,salary) and display record of employees
// having maximum salary.

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

	// Find employee with maximum salary
	maxSalaryIndex := 0
	for i := 1; i < n; i++ {
		if employees[i].Salary > employees[maxSalaryIndex].Salary {
			maxSalaryIndex = i
		}
	}

	// Display employee with maximum salary
	fmt.Println("\nEmployee with Maximum Salary:")
	fmt.Printf("Employee Number: %d\n", employees[maxSalaryIndex].Eno)
	fmt.Printf("Employee Name: %s\n", employees[maxSalaryIndex].Ename)
	fmt.Printf("Salary: %.2f\n", employees[maxSalaryIndex].Salary)
}

