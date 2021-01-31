package concurrency

import (
	"fmt"
	"time"
)

func Example2() {
	fmt.Println("Calling long operation")
	c := make(chan string)
	go LongOperationWithChannel("Test", c)
	fmt.Println("Back in main goroutine")

	// Blocking read from channel
	fmt.Println("Your mutated string ", <-c)
}

func LongOperationWithChannel(s string, c chan string) {
	time.Sleep(time.Second * 2)
	c <- fmt.Sprintf("Your string %v", s)
}
