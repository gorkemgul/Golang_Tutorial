package main

import "fmt"

// Let's create our multiple return function (We need to define the types of the results!)
func mathOperations(number1 int, number2 int) (int, int, int, float32) {
	// Takes two numbers and does a couple of math operations such as: addition, subraction and so on.
	addition := number1 + number2
	subtraction := number1 - number2
	multiplication := number1 * number2
	divison := float32(number1 / number2) // just because the result of divison it's gonna be a decimal number we need to turn it into float

	return addition, subtraction, multiplication, divison
}

func main() {
	// We're gonna explain the concept of multiple return functions by creating a function called mathOperations
	// Let's see how our function works
	result1, result2, result3, result4 := mathOperations(10, 5)
	fmt.Printf("The result of addition operation is: %v", result1)
	fmt.Printf("\nThe result of subtraction operation is: %v", result2)
	fmt.Printf("\nThe result of multiplication operation is: %v", result3)
	fmt.Printf("\nThe result of divison operation is: %v", result4)

	// Note: In case we don't want every operations in our function we could use "_" for operations that we don't want use.
	result5, result6, result7, _ := mathOperations(12, 4)
	fmt.Printf("\nThe result of addition operation is: %v", result5)
	fmt.Printf("\nThe result of subtraction operation is: %v", result6)
	fmt.Printf("\nThe result of multiplication operation is: %v", result7)
	//fmt.Printf("\nThe result of divison operation is: %v", result8)
}
