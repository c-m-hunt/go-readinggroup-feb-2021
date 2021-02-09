package concurrency

import (
	"fmt"
)

// Example3 More of a pipeline type example.
// Two goroutines. One generates numbers 1 to 10 and sends them to
// a channel. The other grabs it, squares it and sends it out
func Example3() {
	cIn := make(chan int, 5)
	cOut := make(chan int)

	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Println("Adding value to in channel ", i)
			cIn <- i
		}
		close(cIn)
	}()

	go func() {
		for {
			fmt.Println("In func but blocked")
			in, ok := <-cIn

			if !ok {
				close(cOut)
				return
			}
			fmt.Println("Got a value!")
			fmt.Println("Received in ", in)
			cOut <- in * in
		}
	}()

	for val := range cOut {
		fmt.Println("Received out ", val)
	}
}
