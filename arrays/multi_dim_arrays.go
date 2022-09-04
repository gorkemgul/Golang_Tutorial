package main

import "fmt"

func main() {
	// Create a multi dimensional array
	var numbers [2][3]int // [rows][columns]

	// Let's change the values of our multi dimensional array
	numbers[0][0] = 1
	numbers[0][1] = 2
	numbers[0][2] = 3
	numbers[1][0] = 4
	numbers[1][1] = 5
	numbers[1][2] = 6

	// Print the array out
	fmt.Println(numbers)

	// Create the for loops to walk through our multi dimensional array
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(numbers[i][j])
		}
	}
}
