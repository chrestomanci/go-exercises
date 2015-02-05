package main

import "fmt"

func main() {
	fmt.Print("Enter a temp in Fahrenheit: ")
	var tempF float64
	fmt.Scanf("%f", &tempF)

	var tempC float64
	tempC = (tempF - 32) * (5.0 / 9.0)

	fmt.Printf("Temp in Centegrade is: %f\n", tempC)
}
