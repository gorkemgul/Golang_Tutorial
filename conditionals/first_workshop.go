package main

import "fmt"

func main() {
	// Task:
	// - Create 3 Integers and find the highest one using conditionals.

	// Creating the integers
	var number1 int = 10
	var number2 int = 25
	var number3 int = 30

	// Create a variable called highestNumber
	var highestNumber int = number1

	// Let's try to find the highest one
	if number2 > highestNumber {
		highestNumber = number2
	}

	if number3 > highestNumber {
		highestNumber = number3
	}

	// Print the highest number out
	fmt.Printf("The highest number is %v", highestNumber)
}
