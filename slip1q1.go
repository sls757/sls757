// B)
// Write a program in GO language to accept n student details like roll_no, [20 Marks]
// stud_name, mark1,mark2, mark3. Calculate the total and average of
// marks using structure.

package main

import "fmt"

type Student struct {
    rollNo int
    name   string
    mark1  float32
    mark2  float32
    mark3  float32
}

func main() {
    var no int
    fmt.Print("\nEnter the number of students: ")
    fmt.Scanln(&no)

    students := make([]Student, no)

    for i := 0; i < no; i++ {
        fmt.Println("\n        Enter details for Student", i+1, "\n")
        fmt.Print("Enter Roll no: ")
        fmt.Scanln(&students[i].rollNo)
        fmt.Print("Enter Student name: ")
        fmt.Scanln(&students[i].name)
        fmt.Print("Enter Mark1: ")
        fmt.Scanln(&students[i].mark1)
        fmt.Print("Enter Mark2: ")
        fmt.Scanln(&students[i].mark2)
        fmt.Print("Enter Mark3: ")
        fmt.Scanln(&students[i].mark3)
    }

    for i, student := range students {
        fmt.Println("\nDetails for Student", i+1)
        fmt.Println("Roll no:", student.rollNo)
        fmt.Println("Name:", student.name)
        fmt.Println("Mark1:", student.mark1)
        fmt.Println("Mark2:", student.mark2)
        fmt.Println("Mark3:", student.mark3)
        calculate(student.mark1, student.mark2, student.mark3)
    }
}

func calculate(mark1, mark2, mark3 float32) {
    total := mark1 + mark2 + mark3
    average := total / 3
    fmt.Println("Total Marks:", total)
    fmt.Println("Average Marks:", average)
}
