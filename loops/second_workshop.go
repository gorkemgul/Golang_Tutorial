package main

import "fmt"

func main() {
	// Task:
	// - We're gonna write a program that hold a number and wants user to find that number

	// Let's start by defining the integers
	var number, predicted_number int
	number = 80          // The number that users are gonna try to predict
	predicted_number = 0 // The number that users are going to enter

	// Create the for loop
	for predicted_number != number {
		// Ask users to enter a number (Ask for a predition)
		fmt.Println("Please enter a number")
		fmt.Scanln(&predicted_number)

		// Create the if-else if-else statement
		if predicted_number < number {
			fmt.Println("Try a higher one!")
		} else if predicted_number > number {
			fmt.Println("Try a lower one!")
		} else {
			fmt.Println("Congratz!, You've found the exact number!")
		}
	}
}
