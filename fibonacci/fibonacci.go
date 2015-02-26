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
		return (fib_no_cache(x-1) + fib_no_cache(x-2))

	}
}

func setup() {
	cache = make(map[uint64]uint64)
	cache[0] = 0
	cache[1] = 1
}

func main() {

	setup()

	fmt.Println("With cache")
	for i := 0; i <= 40; i++ {
		arg := uint64(i)
		fmt.Printf("\tFibonacci %d is %d\n", i, fib(arg))
	}

	fmt.Println("Pure recursive version, Without cache")
	for i := 0; i <= 40; i++ {
		arg := uint64(i)
		fmt.Printf("\tFibonacci %d is %d\n", i, fib_no_cache(arg))
	}

}
