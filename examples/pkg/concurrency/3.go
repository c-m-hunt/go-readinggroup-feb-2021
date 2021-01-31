package concurrency

import (
	"fmt"
)

// Example3 More of a pipeline type example.
// Two goroutines. One generates numbers 1 to 10 and sends them to
// a channel. The other grabs it, squares it and sends it out
func Example3() {
	cIn := make(chan int)
	cOut := make(chan int)

	go func() {
		for {
			fmt.Println("In func but blocked")
			in, ok := <-cIn
			fmt.Println("Got a value!")
			if !ok {
				close(cOut)
				return
			}

			fmt.Println("Received in ", in)
			cOut <- in * in
		}

	}()

	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Println("Adding value to in channel ", i)
			cIn <- i
		}
		close(cIn)
	}()

	for val := range cOut {
		fmt.Println("Received out ", val)
	}

}
