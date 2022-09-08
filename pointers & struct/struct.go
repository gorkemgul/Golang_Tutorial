package main

import "fmt"

func main() {
	// A struct is a user-defined type that allows to group or combine items of possibly different types into a single type
	// Its concept is generally compared with the classes in Object-oriented programming

	// Let's look at how to create a struct
	type product struct {
		name      string
		unitPrice float64
		brand     string
	}

	// Define a couple of products
	product1 := product{"Laptop", 15000, "Monster"}
	product2 := product{"Dishwasher", 8000, "Bosh"}
	product3 := product{"Television", 12000, "Samsung"}
	product4 := product{"Cell Phone", 25000, "Apple"}

	// Let's demonstrate them
	fmt.Printf("The information of the first product is: %v\n", product1)
	fmt.Printf("The information of the second product is: %v\n", product2)
	fmt.Printf("The information of the third product is: %v\n", product3)
	fmt.Printf("The information of the fourth product is: %v\n", product4)

	// In case we don't have all the information then we could use two ways, let's show them
	product5 := product{"Table", 5000, ""}
	product6 := product{name: "Chair", unitPrice: 1000}

	// Let's show their results and compare them to the other ones
	fmt.Printf("The information of the fifth products is: %v\n", product5)
	fmt.Printf("The information of the sixth product is: %v", product6)
}
