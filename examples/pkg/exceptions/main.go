package exceptions

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

func Example1() {
	b, err := ioutil.ReadFile("file/that/does/not/exist")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(b)

}

func Example2() {
	o, err := Divide(6, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(o)

	if o, err = Divide(6, 0); err != nil {
		log.Fatal(err)
	}

	fmt.Println(o)
}

// Divide a by b and return the answer
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Woah, can't divide by zero!")
	}
	return a / b, nil
}
