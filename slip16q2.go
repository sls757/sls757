// B) Write a program in GO language that prints out the numbers from0 
// to 10, waiting between 0 and 250 ms after each one using the
// delay function.

package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		delay(250) // Delay in milliseconds
	}
}

func delay(milliseconds time.Duration) {
	time.Sleep(milliseconds * time.Millisecond)
}
