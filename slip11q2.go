// B) Write a program in GO language to create a buffered channel,
// store few values in it and find channel capacity and length. Read
// values from channel and find modified length of a channel

package main

import (
	"fmt"
)

func main() {
	// Create a buffered channel with a capacity of 3
	bufferedChannel := make(chan int, 3)

	// Store values in the channel
	bufferedChannel <- 10
	bufferedChannel <- 20
	bufferedChannel <- 30

	// Display channel capacity and length
	displayChannelInfo(bufferedChannel, "Initial")

	// Read values from the channel
	value1 := <-bufferedChannel
	value2 := <-bufferedChannel

	// Display modified length after reading values
	displayChannelInfo(bufferedChannel, "Modified")

	fmt.Printf("Read values from the channel: %d, %d\n", value1, value2)

	// Close the channel
	close(bufferedChannel)
}

// Function to display channel capacity and length
func displayChannelInfo(ch chan int, label string) {
	fmt.Printf("%s - Channel Capacity: %d, Length: %d\n", label, cap(ch), len(ch))
}
