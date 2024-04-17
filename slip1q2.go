// Q1. A) Max Marks: 35+15=50
// Write a program in GO language to accept user choice and print answers
// using arithmetic operators.

package main

import "fmt"

func main() {
	var choice int
	var num1 float32
	var num2 float32
	fmt.Println("-----Arithmetic Operator-----")
	fmt.Println("1.Addition")
	fmt.Println("2.Substraction")
	fmt.Println("3.Multiplication")
	fmt.Println("4.Division")
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		getdata(&num1, &num2)
		fmt.Println("Addtion is ", num1, " + ", num2, " = ", num1+num2)
		break
	case 2:
		getdata(&num1, &num2)
		fmt.Println("Substraction is ", num1, " - ", num2, " = ", num1-num2)
		break
	case 3:
		getdata(&num1, &num2)
		fmt.Println("Multiplication is ", num1, " * ", num2, " = ", num1*num2)
		break
	case 4:
		getdata(&num1, &num2)
		if num2 == 0 {
			fmt.Println("can't divide by zero / 0 ! ")
		} else {
			fmt.Println("Division is ", num1, " / ", num2, " = ", num1/num2)
		}
		break
	default:
		fmt.Println("Please enter valid choice!(1,2,3,4)")
		break

	}
}

func getdata(num1, num2 *float32) {
	fmt.Print("\nEnter 1st Number: ")
	fmt.Scanln(num1)
	fmt.Print("Enter 2nd Number: ")
	fmt.Scanln(num2)
}
