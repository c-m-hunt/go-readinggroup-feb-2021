package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func GeneratePrimes(noPrimes int) []int64 {
	primes := []int64{2}

	for n := int64(3); n > 0; n += 2 {
		prime := true

		r := int64(math.Sqrt(float64(n))) + 1
		for i := int64(3); i < r; i += 2 {
			if n%i == 0 {
				prime = false
				break
			}
		}

		if prime {
			primes = append(primes, n)
			if len(primes) == noPrimes {
				return primes
			}
		}
	}
	return primes
}

func main() {
	noPrimes, _ := strconv.Atoi(os.Args[1])

	primes := GeneratePrimes(noPrimes)

	fmt.Printf("%v primes generated", len(primes))
}
