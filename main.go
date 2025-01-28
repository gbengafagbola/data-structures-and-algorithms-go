package main

import (
	"fmt"
	"math/rand"
	"time"
)

// sleep function
func sleep(){
	time.Sleep(
		time.Duration(rand.Intn(3000)) *  time.Millisecond,
	)
}

// producer: which write message to a channel
// write only 
func producer(ch chan <- int, name string) {
	for {
		// sleep for some random time
		sleep()
		// generate a random number
		n := rand.Intn(100)

		//send the message 
		fmt.Printf("Channel %s -> %d\n", name, n)

	}
}

// consumer: read the message
// read only
func consumer(ch <- chan int){
	for n := range ch {
		fmt.Printf("<- %d\n", n)
	}
}

// basically this will read from channel A & B and write their content into channel C
func fanIn(chA, chB <- chan int, chC chan <- int){
	var n int 
	for {
		select {
		case n = <- chA:
			chC <- n
		case n = <- chB:
			chC <- n
		}
	}
}


func main (){
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go producer(chA, "A")
	go producer(chB, "B")
	go consumer(chC)

	fanIn(chA, chB, chC)
}