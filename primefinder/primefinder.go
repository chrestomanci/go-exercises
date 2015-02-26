package main

import (
	"fmt"
)

// TODO: Replace with a list
var primes map[uint]bool

func test(num uint) bool {

	for prime := range primes {

		if 2 == prime {
			continue
		}

//		fmt.Printf("\t\tTesting %d with %d\n", num, prime)

		if num == prime {
			return true
		}

		if 0 == num%prime {
			return false
		}
	}

	return true
}

func main() {

	setup()

	var i uint

	for i = 11; i <= 10000; i += 2 {
		if test(i) {
			fmt.Println("Is prime: ", i)
			primes[i] = true
		} else {
			fmt.Println("\tNot prime: ", i)
		}

	}
}

func setup() {
	primes = make(map[uint]bool)
	primes[2] = true
	primes[3] = true
	primes[5] = true
	primes[7] = true
	primes[11] = true
	primes[13] = true
	primes[17] = true
	primes[19] = true
}
