package main

import (
	"fmt"
	"github.com/chrestomanci/go-exercises/listutil"
)

func main() {
	list := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	min := listutil.Min(list)

	fmt.Printf("The smalles number from %v is: %d\n", list, min)
}
