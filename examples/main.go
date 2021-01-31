package main

import (
	"os"

	"github.com/c-m-hunt/go-readinggroup/pkg/concurrency"
	"github.com/c-m-hunt/go-readinggroup/pkg/exceptions"
)

func main() {
	example := os.Args[1]

	switch example {
	case "1":
		exceptions.Example1()
	case "2":
		exceptions.Example2()
	case "3":
		concurrency.Example1()
	case "4":
		concurrency.Example2()
	case "5":
		concurrency.Example3()
	}

}
