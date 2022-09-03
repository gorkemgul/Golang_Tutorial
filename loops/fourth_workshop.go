package main

import "fmt"

func main() {
	// Task:
	// - We're gonna write a program that decides if given numbers are friendly numbers or not

	// Let's start by defining variables
	var number1, number2, result1, result2 int
	number1, number2 = 220, 284
	result1, result2 = 0, 0

	// Let's create a for loop for number1
	for i := 1; i < number1; i++ {
		if number1%i == 0 {
			result1 += number1
		}
	}

	// Let's create another for loop for number2
	for i := 1; i < number2; i++ {
		if number2%i == 0 {
			result2 += number2
		}
	}

	// Let's check if they're friendly numbers or not
	if result1 == number2 && result2 == number1 {
		fmt.Printf("%v and %v are friendly numbers.", number1, number2)
	} else {
		fmt.Printf("%v and %v are not friendly numbers.", number1, number2)
	}
}
