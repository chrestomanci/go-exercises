package main

import (
	"fmt"
)

// Setup an initial slice with a seed list
var primes = []uint{
	2,
	3,
	5,
	7,
	11,
	13,
	17,
	19,
}

func test(num uint) bool {

	for _, prime := range primes {

		if 2 == prime {
			continue // No need to test, caller will only send odd numbers.
		}

		if 0 == prime {
			// ran out of entries in the slice
			return true
		}

		// fmt.Printf("\t\tTesting %d with %d\n", num, prime)

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

	var i uint

	for i = 21; i <= 1000; i += 2 {
		if test(i) {
			fmt.Println("Is prime: ", i)
			primes = append(primes, i)
		} else {
			// fmt.Println("\tNot prime: ", i)
		}

	}
}
