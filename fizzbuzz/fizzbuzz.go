package main

import "fmt"

func main() {
	//	by_loop()
	by_map()
}

func by_loop() {
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

func by_map() {
	numbers := make(map[int]string)

	for i := 0; i <= 100; i += 3 {
		numbers[i] += "Fizz"
	}

	for i := 0; i <= 100; i += 5 {
		numbers[i] += "Buzz"
	}

	for i := 1; i <= 100; i++ {
		content, ok := numbers[i]

		if ok {
			fmt.Println(content)
		} else {
			fmt.Println(i)
		}
	}
}
