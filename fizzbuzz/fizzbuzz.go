package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if 0 != i%3 && 0 != i%5 {
			fmt.Println(i)
		} else {
			if 0 == i%3 {
				fmt.Print("fizz")
			}

			if 0 == i%5 {
				fmt.Print("buzz")
			}

			fmt.Println("") // print newline
		}
	}
}
