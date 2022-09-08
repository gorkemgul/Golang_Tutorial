package main

import "fmt"

// Here, we're gonna work with pointers which are very important for memory management
// Pointers is a variable that is used to store the memory address of another variable.
// The reason why we're using pointers is that remembering all the memory addresses manually is overhead.
// We must use two parameters in pointers:
// - "*" operator: we use the dereferencing operator to declare pointer variable and access the value stored in the address.
// - "&" operator: we use the address operator to return the address of a variable or to access the address of a variable to a pointer.

// Let's create a function to have a better understanding of the concept of the pointers
func addition(number int) {
	number = number + 1
	fmt.Printf("The value of number is %v\n", number)
}

// Re-create our addition function
func addition2(number *int) {
	*number = *number + 1
	fmt.Printf("The value of number is %v\n", *number)
}

// Define a function for our array
func addition3(array []int) {
	array[1] = 100
	fmt.Printf("The result of the array is: %v", array)
}

func main() {
	// Let's create a number and use it in our function
	var number int = 20

	// Use the addition function and compare the results
	addition(number)
	fmt.Printf("The value of our number variable is %v\n", number)
	// Note: As we've seen, the value of our number variable and the output of our function is different. If we want them to be the same we need to use pointers.

	// Let's create another number called number2 but this time we're gonna use pointers
	var number2 int = 25

	// Let's use the addition2 function this time
	addition2(&number2)
	fmt.Printf("The value of our number2 variable is %v\n", number2)
	// Note: As we've seen they're equal now.

	// Let's make another example to have a better understanding of the concept
	// Create two variables called a and b
	var a int = 30
	var b *int = &a

	// Demonstrate the results
	fmt.Printf("The value of a is equal to %v\n", a)
	fmt.Printf("The address of a is equal to %v\n", &a)
	fmt.Printf("The value stored in b is: %v\n", *b)

	// When it comes to arrays, they work with the addresses (if we make a change on one of its indexes then the original one is also going to change!)
	originalArray := []int{1, 2, 3, 4, 5}

	// Use the addition3 function
	addition3(originalArray)
	fmt.Printf("\nOriginal Array: %v", originalArray)
}
