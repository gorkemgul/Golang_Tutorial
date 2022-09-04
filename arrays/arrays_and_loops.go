package main

import "fmt"

func main() {
	// Create an array of 5 countries
	var countries [5]string
	countries[0] = "Turkey"
	countries[1] = "France"
	countries[2] = "UK"
	countries[3] = "USA"
	countries[4] = "Germany"

	// Let's print our countries array out
	fmt.Println(countries)

	// To walk through a list, we generally use a for loop
	for i := 0; i < 5; i++ {
		fmt.Println(countries[i])
	}

}
