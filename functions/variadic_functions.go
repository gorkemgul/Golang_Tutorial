package main

import "fmt"

// We're gonna explain the concept of variadic functions here
// Variadic functions can take infinite parameters which is why they're very useful in programming.
// Let's explain shortly when we use variadic functions:
// - They can be used instead of passing a slice to a function
// - They can be used when we don’t know the number of parameters
// - They're especially useful when the arguments to our function are most likely not going to come from a single data structure (i.e. slice or array etc.)

// Let's start by creating a variadic function
func addition(numbers ...int) int {
	// Takes numbers and turns them into an array, lastly does an addition operation and returns an output
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func main() {
	// Let's use our function called addition
	result := addition(1, 2, 3, 4, 5)
	fmt.Println(result)

	// In case we don't enter any number into our function it's still gonna work and return 0 as a result
	result2 := addition()
	fmt.Println(result2)

	// We could also create an array and pass it to our function
	numbers := []int{10, 20, 30, 40, 50}
	result3 := addition(numbers...) // Note: When we want to enter an array into our variadic function we need to enter it like this way, otherwise it's not gonna work.
	fmt.Println(result3)
}
