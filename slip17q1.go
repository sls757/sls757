// Q1. A) Write a program in GO language to illustrate the concept of
// returning multiple values from a function. ( Add, Subtract,
// Multiply, Divide)

package main

import "fmt"

func performArithmeticOperations(a, b float64) (float64, float64, float64, float64) {
    // Add
    addResult := a + b

    // Subtract
    subtractResult := a - b

    // Multiply
    multiplyResult := a * b

    // Divide (check for division by zero)
    var divideResult float64
    if b != 0 {
        divideResult = a / b
    }

    return addResult, subtractResult, multiplyResult, divideResult
}

func main() {
    // Example values
    num1 := 10.0
    num2 := 5.0

    // Calling the function and receiving multiple values
    add, subtract, multiply, divide := performArithmeticOperations(num1, num2)

    // Displaying the results
    fmt.Printf("Addition: %.2f\n", add)
    fmt.Printf("Subtraction: %.2f\n", subtract)
    fmt.Printf("Multiplication: %.2f\n", multiply)

    if num2 != 0 {
        fmt.Printf("Division: %.2f\n", divide)
    } else {
        fmt.Println("Cannot divide by zero.")
    }
}
