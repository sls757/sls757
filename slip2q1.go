// Q1. A)
// Max Marks: 35+15=50
// Write a program in GO language to print Fibonacci series of n
// terms.
package main

import "fmt"

func main() {
    var no int
    fmt.Print("\nEnter the number of terms for Fibonacci series: ")
    fmt.Scanln(&no)

    a, b := 0, 1
    fmt.Print("\nFibonacci Series:")
    for i := 0; i < no; i++ {
        fmt.Print(" ", a)
         a, b = b, a+b
    }
}
