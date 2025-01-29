package main

import (
	"fmt"
	"math/rand"
	"time"
)

// sleep function introduces a random delay to simulate processing time
func sleep() {
	time.Sleep(
		time.Duration(rand.Intn(3000)) * time.Millisecond, // Sleep for a random duration between 0 and 3000 milliseconds
	)
}

// producer generates random numbers and writes them to a channel
// ch is a write-only channel (chan <- int)
func producer(ch chan<- int) {
	for {
		// Sleep for some random time before producing a value
		sleep()

		// Generate a random number between 0 and 99
		n := rand.Intn(100)

		// Print the generated value and send it to the channel
		fmt.Printf(" -> %d\n", n)
		ch <- n
	}
}

// consumer reads values from a channel and processes them
// ch is a read-only channel (<- chan int)
// name is used to identify the consumer in the output
func consumer(ch <-chan int, name string) {
	for n := range ch {
		// Print the value read from the channel with the consumer's name
		fmt.Printf("Consumer %s <- %d\n", name, n)
	}
}

// fanOut function distributes values from one input channel (chA)
// to two output channels (chB and chC) based on a condition
func fanOut(chA <-chan int, chB, chC chan<- int) {
	for n := range chA { // Continuously read from input channel chA
		if n < 50 {
			// If the value is less than 50, send it to chB
			chB <- n
		} else {
			// Otherwise, send it to chC
			chC <- n
		}
	}
}

func main() {
	// Create three channels:
	// chA is the input channel for the producer
	// chB and chC are the output channels for the consumers
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	// Start the producer goroutine that writes values to chA
	go producer(chA)

	// Start two consumer goroutines:
	// Consumer "B" reads from chB
	go consumer(chB, "B")
	// Consumer "C" reads from chC
	go consumer(chC, "C")

	// Call fanOut to distribute values from chA to chB and chC
	fanOut(chA, chB, chC)
}
