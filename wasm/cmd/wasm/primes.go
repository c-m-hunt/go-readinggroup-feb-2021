package main

import (
	"fmt"
	"math"
	"syscall/js"
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

func primesWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		noPrimes := args[0].Int()
		fmt.Printf("input %v\n", noPrimes)

		primes := GeneratePrimes(noPrimes)
		primesOut := make([]interface{}, len(primes))
		for i, p := range primes {
			primesOut[i] = p
		}

		return primesOut
	})
	return jsonFunc
}
