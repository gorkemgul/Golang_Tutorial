package main

import "fmt"

func main() {
	// Let's create a slice using another way
	countries := []string{"Turkey", "France", "Germany", "UK", "Canada"}

	// Demonstrate the slice
	fmt.Printf("The original Slice: %v", countries)

	// Create an empty slice using the same size of countries slice
	copiedCountries := make([]string, len(countries))

	// Demonstrate the copiedCountries slice
	fmt.Printf("\nBefore using the copy function: %v", copiedCountries)

	// Let's copy the countries slice
	copy(copiedCountries, countries)

	// Visualize the copiedCountries slice after using copy func
	fmt.Printf("\nAfter using the copy function: %v", copiedCountries)

	// Let's explore what would happen if we add another variable into the original slice
	countries = append(countries, "USA")

	// Print countries slice out after adding another country
	fmt.Printf("\nAfter adding another country: %v", countries)

	// Did the copied slice also change?
	fmt.Printf("\ncopiedCountries: %v\n", copiedCountries)
	// Note: It doesn't matter if we change the original slice (e.g. adding another value and etc.) which we copied the values of its and created another one from it. The copied one is gonna keep the same elements in it.

	// Let's look at how to do slicing operations
	fmt.Println(countries[1:3]) // get the values from 1 to 3 (not including 3!)
	fmt.Println(countries[1:5]) // get the values from 1 to 5 (again, not including 5!)
}
