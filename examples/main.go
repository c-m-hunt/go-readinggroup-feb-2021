package main

import (
	"os"

	"github.com/c-m-hunt/go-readinggroup/pkg/concurrency"
	"github.com/c-m-hunt/go-readinggroup/pkg/exceptions"
	"github.com/c-m-hunt/go-readinggroup/pkg/project"
)

func main() {
	example := os.Args[1]

	switch example {
	case "ex1":
		exceptions.Example1()
	case "ex2":
		exceptions.Example2()
	case "con1":
		concurrency.Example1()
	case "con2":
		concurrency.Example2()
	case "con3":
		concurrency.Example3()
	case "run":
		project.Run()
	}

}
