package concurrency

import (
	"fmt"
	"time"
)

func Example1() {
	fmt.Println("Calling long operation")
	go LongOperation("Test")
	fmt.Println("Back in main goroutine")

	// Wait for input
	fmt.Scanln()
}

func LongOperation(s string) {
	time.Sleep(time.Second * 2)
	fmt.Println("Your string ", s)
}
