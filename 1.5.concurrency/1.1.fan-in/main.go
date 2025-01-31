package main

import (
	"fmt"
	"math/rand"
	"time"
)

// sleep function introduces a random delay to simulate work or waiting
func sleep() {
	time.Sleep(
		time.Duration(rand.Intn(3000)) * time.Millisecond, // Sleep for a random duration between 0 and 3000 milliseconds
	)
} 

// producer function generates random numbers and sends them to a channel
// ch is a write-only channel (chan <- int)
// name identifies the producer (useful for distinguishing producers)
func producer(ch chan<- int, name string) {
	for {
		// Sleep for a random time
		sleep()

		// Generate a random number
		n := rand.Intn(100)

		// Send the random number to the channel
		// Print a message indicating which producer sent the value
		fmt.Printf("Channel %s -> %d\n", name, n)

		// Send the value to the channel
		ch <- n
	}
}

// consumer function continuously reads numbers from a channel and prints them
// ch is a read-only channel (<- chan int)
func consumer(ch <-chan int) {
	for n := range ch {
		// Print the received value
		fmt.Printf("<- %d\n", n)
	}
}

// fanIn function reads from two input channels (chA and chB)
// and forwards their values to a single output channel (chC)
func fanIn(chA, chB <-chan int, chC chan<- int) {
	var n int
	for {
		// Use select to listen to multiple channels simultaneously
		select {
		case n = <-chA:
			// When a value is received from chA, forward it to chC
			chC <- n
		case n = <-chB:
			// When a value is received from chB, forward it to chC
			chC <- n
		}
	}
}

func main() {
	// Create three channels:
	// chA and chB for producers
	// chC for the final output after merging the two producer channels
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	// Start producer goroutines
	// Producer A sends values to chA
	go producer(chA, "A")
	// Producer B sends values to chB
	go producer(chB, "B")

	// Start a consumer goroutine that reads from chC
	go consumer(chC)

	// fanIn function merges values from chA and chB into chC
	fanIn(chA, chB, chC)
}
