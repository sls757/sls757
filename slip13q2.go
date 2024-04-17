// B) Write a function in GO language to find the square of a number
// and write a benchmark for it.

package main

import (
	"fmt"
	"testing"
)

// Function to find the square of a number
func square(num int) int {
	return num * num
}

// Benchmark function for square
func BenchmarkSquare(b *testing.B) {
	// Run the square function b.N times
	for i := 0; i < b.N; i++ {
		square(5) // You can change the input value as needed
	}
}

func main() {
	// Example usage of the square function
	result := square(3)
	fmt.Printf("Square of 3: %d\n", result)

	// Run the benchmark
	fmt.Println("Running benchmark for square function...")
	resultBenchmark := testing.Benchmark(BenchmarkSquare)
	fmt.Printf("Benchmark result: %s\n", resultBenchmark)
}
