package main

import (
	"fmt"
	"time"
)

// Goroutines is a function or method which executes independently and simultaneously in connection with any other Goroutines present in our program
// We can consider a goroutine like a light weighted thread.

// Let's explain the concept a little bit more by creating two function and using them with and without goroutines
// Create the first function to get even numbers between 0 and given number
func getEven(number int) {
	for i := 0; i < number; i += 2 {
		fmt.Printf("%v is a even number\n", i)
		// time.Sleep(1 * time.Second) (optional to display better)
	}
}

// Create the second function to get odd numbers between 1 and given number
func getOdd(number int) {
	for i := 1; i < number; i += 2 {
		fmt.Printf("%v is an odd number\n", i)
		// time.Sleep(1 * time.Second) (optional to display better)
	}
}

func main() {
	// Let's create an integer called number and use our functions
	var number int = 10
	getEven(number)
	getOdd(number)
	// Note: As we've seen our algorithm worked from top to end step by step, but we want our to get results at the same time. (simultaneously) to do so, let's use goroutine

	go getEven(number)
	go getOdd(number)
	time.Sleep(5 * time.Second)
	// Note: Now, as we've seen in the result our functions work simultaneously, to display them better we may add time to our functions
}
