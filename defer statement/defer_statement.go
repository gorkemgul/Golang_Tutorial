package main

import "fmt"

// Defer statements delay the execution of the function or method until the nearby functions returns.
// Let's define a few functions to explain the concept of the defer statement better.
func Hello() {
	fmt.Println("Hello")
}

func World() {
	fmt.Println("World")
}

func helloWorld() {
	fmt.Println("Hello World!")
}

func main() {
	defer World()
	defer Hello()
	helloWorld()
	// Note: Multiple defer statements are allowed in the same program and they are executed in LIFO (Last-In, First-Out).
}
