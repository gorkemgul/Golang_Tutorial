package main

import "fmt"

func main() {
	// The reason why we use a slice is, an array has a fixed size, on the other hand A slice is a dynamically-sized, flexible view into the elements of an array
	// Create a slice
	names := make([]string, 3)

	// Print it out
	fmt.Printf("Before adding the values: %v", names)

	// Let's define a couple of values to our slice
	names[0] = "Gorkem"
	names[1] = "Recep"
	names[2] = "Caleb"

	// To see the difference let's check it again
	fmt.Printf("\nAfter adding the values: %v", names)

	// We could add another value to our slice by using append function
	names = append(names, "Mark")

	// Check it againg after adding a new name
	fmt.Printf("\n After adding a name with append func: %v", names)
}
