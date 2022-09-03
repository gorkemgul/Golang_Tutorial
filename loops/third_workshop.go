package main

import "fmt"

func main() {
	// Task:
	// - Write a program that decides if the number which is taken from user is Prime number or not
	// Note: A prime number is a whole number greater than 1 whose only factors are 1 and itself

	// Let's start by creating an integer and get a number from user
	number := 0
	fmt.Println("Please enter a number!")
	fmt.Scanln(&number)

	// Create a boolean variable called prime_number
	var prime_number bool = true // it's gonna be true by default

	// Create a for loop to check if user's number is prime or not
	for i := 2; i < number; i++ {
		if number%i == 0 {
			prime_number = false
		}
	}
	if prime_number == true {
		fmt.Println("The number you've entered is a prime number")
	} else {
		fmt.Println("The number you've entered is not a prime number!")
	}
}
