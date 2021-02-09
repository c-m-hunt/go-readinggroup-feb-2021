package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("generatePrimes", primesWrapper())
	<-make(chan bool)
}
