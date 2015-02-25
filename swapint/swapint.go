package main

import (
	"fmt"
)

func main() {

	a := 1
	b := 2

	swapint(&a, &b)

	fmt.Printf("a is now %d, b is now %d\n", a, b)

}

func swapint(x *int, y *int) {
	var tmp int

	tmp = *x
	*x = *y
	*y = tmp

	return
}

func square(x *float64) {
	*x = *x * *x
}
