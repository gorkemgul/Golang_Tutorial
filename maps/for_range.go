package main

import "fmt"

func main() {
	// We're gonna cover how to create a for loop with range syntax
	// We use for-range to walk through an array, a map and etc.
	// Let's create an array
	fruits := []string{"Apple", "Banana", "Orange", "Watermelon", "Melon"}

	// Use only for loop to walk through our array
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	// Let's look at how to walk through using for-range
	for i, fruit := range fruits {
		fmt.Printf("The index of the element is: %v, The element is: %v\n", i, fruit)
	}

	// In case we don't want to get indexes, we can use "_" instead of writing "i"
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// Lastly, let's create a map and walk through it
	office_stuff := map[string]string{"1st": "Pencil", "2nd": "Paper", "3rd": "Printer", "4th": "Scissors"}
	for k, v := range office_stuff { // k stands for key & v stands for value
		fmt.Printf("The key: %v, The value: %v\n", k, v)
	}
}
