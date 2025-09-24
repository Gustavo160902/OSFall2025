// Gustavo Lujan - 09/24/2025
// HW1 - Producer-Consumer with Goroutines and Channels in Go
// recerating the hw#0 behaivor

package main

import (
	"fmt"	
)

// Producer function to send numbers to the channel.Only one that uses the if stamtement.
func producer(ch chan int, done chan bool) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producer: %d\n", i)
		ch <- i // send to consumer
		<-done  // wait for consumer to finish
	}
	close(ch) // close channel when done
}

// Consumer function to receive numbers from the channel and print them.
func consumer(ch chan int, done chan bool) {
	for val := range ch { // receive until channel closed
		fmt.Printf("Consumer: %d\n", val)
		done <- true // notify producer to continue IA help I did not know how to return it to the producer function.
	}
}

func main() {
	ch := make(chan int)   
	done := make(chan bool) 

	go consumer(ch, done)
	producer(ch, done)
}

