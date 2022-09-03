package main

import "fmt"

func main() {
	// GO language contains only a single loop that is for-loop.
	// Let's say we want to print a text for five times. We could print the text by using print method for 5 times, but this way couldn't be efficient.
	// Let's use for loop to print the text for five times.

	// Create an integer and a string
	i := 0
	hello_world := "Hello World"
	// Create the for loop
	for i <= 4 {
		fmt.Println(hello_world)
		i += 1
	} // Note: The reason why we added 1 each time is if we didn't our for loop could've became an infinite for loop

	// Show an example of infinite loop
	for i <= 4 {
		fmt.Println(hello_world)
	} // If we run this one we're gonna get an infinite loop
}
