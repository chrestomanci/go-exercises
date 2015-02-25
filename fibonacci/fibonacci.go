package main

import (
	"fmt"
)

var cache map[uint64]uint64

func fib(x uint64) uint64 {

	cached, ok := cache[x]

	if ok {
		return cached
	} else {
		result := (fib(x-1) + fib(x-2))
		cache[x] = result
		return result
	}
}

// About 10 times slower
func fib_no_cache(x uint64) uint64 {
	switch x {
	case 0:
		return 0

	case 1:
		return 1

	default:
		return (fib(x-1) + fib(x-2))

	}
}

func setup() {
	cache = make(map[uint64]uint64)
	cache[0] = 0
	cache[1] = 1
}

func main() {

	setup()

	for i := 1; i <= 80; i++ {
		arg := uint64(i)
		fmt.Printf("Fibonacci %d is %d\n", i, fib(arg))
	}

}
