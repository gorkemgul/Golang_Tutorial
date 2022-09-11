package main

import "fmt"

// We're gonna do a basic example to show something important because we need to be aware while using defer statement
// Let's create a function that detects odd & even numbers
func oddeven(number int) string {
	// Let's add another function with defer statement. Here our goal is to show what if we add the defer statement to the top and to the end of the function. Let's find out!
	// Add the defer statement to the top
	defer helloWorld()
	if number%2 == 0 {
		fmt.Println("If statement is worked!") // we've written this, to prove that if statements works before the defer statement
		return "This is an even number!"
	}

	if number%2 != 0 {
		return "This is an odd number!"
	}
	// Add the defer statement to the end
	// defer helloWorld()
	// Note: As we've seen if we write it to the end it's not gonna work because if one of its if statements is True and returns an output the functions ends.
	return "Undefined!"
}

// Create a function to print hello world
func helloWorld() {
	fmt.Println("Hello World!")
}

func main() {
	// Let's test our function
	result := oddeven(4)
	fmt.Println(result)
}
