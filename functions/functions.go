package main

import "fmt"

// We're gonna explain how to create a couple of functions and show how to use them
// Let's create a function for addition operations
func summation(number1 int, number2 int) int {
	// Takes two numbers to make an addition and prints the result out
	result := number1 + number2
	return result
}

// Create a function for subtraction operations
func subtraction(number1 int, number2 int) int {
	// Takes two numbers to make a subtraction and prints the result out
	result := number1 - number2
	return result
}

// Create a function to print Hello World
func helloWorld() {
	fmt.Println("Hello World!")
}

// Let's use our defined functions
func main() {
	// Test summation function
	result := summation(10, 5)
	fmt.Println(result)

	// Test subtraction function
	result2 := subtraction(10, 5)
	fmt.Println(result2)

	// Test helloWorld function
	helloWorld()
}
