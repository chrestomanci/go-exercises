package main

import "fmt"

func main() {
	fmt.Print("Enter an even number: ")
	var before int
	fmt.Scanf("%d", &before)

	half, ok := half(before)

	if ok {
		fmt.Printf("Half of %d is %d\n", before, half)
	} else {
		fmt.Printf("%d is not an even number\n", before)
	}
}

func half(before int) (int, bool) {
	if 0 == before%2 {
		half := before / 2
		return half, true
	} else {
		return 0, false
	}
}
