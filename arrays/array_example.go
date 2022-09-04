package main

import "fmt"

func main() {
	// We're gonna do an exercise for practicing
	// We're gonna write a program that finds out which number is the highest one in the given array

	// Let's start by creating an array called numbers
	numbers := [5]int{20, 30, 50, 10, 2}

	// Create an integer called highestNumber
	highestNumber := 0

	// Create the for loop to walk through our list
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > highestNumber {
			highestNumber = numbers[i]
		}
	}
	// Let's check the highest number
	fmt.Printf("The highest number is %v", highestNumber)
}
