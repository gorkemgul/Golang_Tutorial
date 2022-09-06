package main

import "fmt"

func main() {
	// Maps is a collection of unordered pairs of key-value. It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys (like dictionary key-value)
	// Let's create one to work with
	animals := make(map[string]string)

	// Let's add values and keys into it
	animals["First"] = "Lion"
	animals["Second"] = "Parrot"
	animals["Third"] = "Cat"
	animals["Fourth"] = "Dog"
	animals["Fifth"] = "Fish"
	animals["Sixth"] = "Bear"

	// Print it out
	fmt.Println(animals)

	// We could check how many parameters are there in our map
	numberParameters := len(animals)
	fmt.Printf("\nThere are %v parameters in our map", numberParameters)

	// We can delete a parameter from our map
	delete(animals, "Fourth")

	// Let's check our map and its parameter's number after deleting one parameter
	fmt.Println("\nOur map after deleting one parameter:", animals)
	fmt.Printf("\nThe length of our map after deleting one parameter is: %v", len(animals))

	// Check a specific value and if it exists in map
	value, exist := animals["Third"]
	fmt.Printf("\nValue: %v\nIs it in the map: %v", value, exist)

	// Check the parameter we've deleted from the map
	value2, exist2 := animals["Fourth"]
	fmt.Printf("\nValue: %v\nIs it in the map: %v\n", value2, exist2)

	// There is another way to create a map. Let's look at it (without make() function)
	animals2 := map[string]string{"first": "Panda", "Second": "Snake"}

	// Let's demonstrate the map called animals2
	fmt.Println(animals2)
}
