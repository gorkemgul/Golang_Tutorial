package main

import "fmt"

func main() {
	// Create an array called numbers
	var numbers [5]int

	// Let's print our array out
	fmt.Println("Before we change the value of the second index:", numbers)
	// Note: If we define an array like above, we get an array of zeros ([0 0 0 0 0]) (5 zeros because we've created it with 5 elements)

	// Let's change the value of a specific array
	numbers[4] = 12

	// Print our array out again
	fmt.Println("After we changed the value of the second index:", numbers)
}
