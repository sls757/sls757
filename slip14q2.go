// B) Write a program in GO language using go routine and channel that
// will print the sum of the squares and cubes of the individual digits
// of a number. Example if number is 123 then
// squares = (1 * 1) + (2 * 2) + (3 * 3)
// cubes = (1 * 1 * 1) + (2 * 2 * 2) + (3 * 3 * 3).

package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	// Enter the number for which you want to calculate the sum of squares and cubes
	number := 123

	// Create channels to store results
	squaresCh := make(chan int)
	cubesCh := make(chan int)

	// Use WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Start goroutines to calculate squares and cubes
	wg.Add(2)
	go calculateSquares(number, squaresCh, &wg)
	go calculateCubes(number, cubesCh, &wg)

	// Use a separate WaitGroup to wait for closing channels after goroutines finish
	var closeWG sync.WaitGroup
	closeWG.Add(2)

	// Wait for both goroutines to finish
	go func() {
		wg.Wait()
		close(squaresCh)
		closeWG.Done()
	}()

	go func() {
		closeWG.Wait()
		close(cubesCh)
	}()

	// Retrieve results from channels
	squares := <-squaresCh
	cubes := <-cubesCh

	// Print the results
	fmt.Printf("Sum of squares: %d\n", squares)
	fmt.Printf("Sum of cubes: %d\n", cubes)
}

// Function to calculate the sum of squares of individual digits
func calculateSquares(num int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	digits := getDigits(num)
	sum := 0

	for _, digit := range digits {
		sum += digit * digit
	}

	ch <- sum
}

// Function to calculate the sum of cubes of individual digits
func calculateCubes(num int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	digits := getDigits(num)
	sum := 0

	for _, digit := range digits {
		sum += digit * digit * digit
	}

	ch <- sum
}

// Function to get individual digits of a number
func getDigits(num int) []int {
	strNum := strconv.Itoa(num)
	digits := make([]int, len(strNum))

	for i, char := range strNum {
		digit, _ := strconv.Atoi(string(char))
		digits[i] = digit
	}

	return digits
}


