package main

import "fmt"

// A channel is a medium through which a goroutine communicates with another goroutine and this communication is lock-free. Or in other words, a channel is a technique which allows to let one goroutine to send data to another goroutine
// Let's explain the concept a little bit more by creating two function and using them
// Create the first function to get even numbers between 0 and 10 (including 10)
func getEven(evenNumberCn chan int) {
	result := 0
	for i := 0; i <= 10; i += 2 {
		result += i
	}

	// sending operation
	evenNumberCn <- result
}

// Create the second function to get odd numbers between 1 and 10 (including 10)
func getOdd(oddNumberCn chan int) {
	result := 0
	for i := 1; i <= 10; i += 2 {
		result += i
	}

	// sending operation
	oddNumberCn <- result
}

func main() {

	// Let's create the channels
	evenNumCn := make(chan int)
	oddNumCn := make(chan int)

	// Let's use our functions and see if our channels work or not
	go getEven(evenNumCn)
	go getOdd(oddNumCn)
	// time.Sleep(5 * time.Second) // Note: we don't need to set time while working with channels

	// In case we want to do an operation with the results of our functions we need to receive the data to do so, let's create the result variables
	result1 := <-evenNumCn
	result2 := <-oddNumCn

	// Let's multiply the results and print it out
	result := result1 * result2
	fmt.Printf("The result of multiplication is %v", result)
}
