// Gustavo Lujan - 09/16/2025
// Note: AI helped me understand how to connect the input/output pipes using os and os/exec
// between producer and consumer. It also helped brainstorm using the Go standard libraries
// "os" and "os/exec" to make two processes talk to each other with pipes.

package main

import (
	"fmt"
	"os"
	"os/exec"
)

// the main process
func producer() {
	child := exec.Command(os.Args[0], "consumer") // create consumer process

	sendPipe, _ := child.StdinPipe()  // pipe that consumer reads from
	child.Stdout = os.Stdout          // connect consumer stdout to terminal
	recvPipe, _ := child.StderrPipe() // pipe that consumer writes to

	child.Start() //starts the consumer process

	buf := make([]byte, 10) // buffer just to hold ack messages

	for i := 1; i <= 5; i++ {
		fmt.Printf("Producer: %d\n", i)
		fmt.Fprintln(sendPipe, i) // send number to consumer
		recvPipe.Read(buf)        // wait for ack from consumer
	}
	child.Wait() // wait until consumer finishes
}

// the second process
func consumer() {
	var num int
	for i := 0; i < 5; i++ {
		fmt.Fscan(os.Stdin, &num) // read from producer
		fmt.Printf("Consumer: %d\n", num)
		fmt.Fprintln(os.Stderr, "ack") // send ack to producer
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "consumer" {
		consumer()
	} else {
		producer()
	}
}
