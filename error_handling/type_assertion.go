package main

import "fmt"

// Create a function to check and verify if the type of given number is integer
func verify(i interface{}) {
	number, ok := i.(int) // if the given number is integer then ok will be equal to "True", and number will be equal to given number. Otherwise ok will be equal to False and number will be equal to 0 (default)
	fmt.Printf("Given number: %v, The type of it is: %v\n", number, ok)
}

func main() {
	// Let's test our function
	verify(5)
	var number1 interface{} = 12
	verify(number1)

	// Let's test it with a string and see the result
	var text interface{} = "Gorkem"
	verify(text)
}
