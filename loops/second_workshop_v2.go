package main

import "fmt"

func main() {
	// Tasks: (Same as before but needs an update to count how many predictions did users make)
	// - We're gonna write a program that hold a number and wants user to find that number

	// Let's start by defining the integers
	var number, predicted_number, number_of_predictions int
	number = 80               // The number that users are gonna try to predict
	predicted_number = 0      // The number that users are going to enter
	number_of_predictions = 0 // The number of predictions

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
		// Count how many predictions have done so far
		number_of_predictions += 1
	}
	fmt.Printf("You've done %v predictions(s)", number_of_predictions)
}
