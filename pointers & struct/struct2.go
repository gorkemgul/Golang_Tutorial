package main

import "fmt"

// Here, we're gonna look at how to add a method to struct type
// Define a struct
type customer struct {
	firstName string
	lastName  string
	age       int
	gender    string
}

// Define a function
func (c customer) save() {
	fmt.Println("The customer was saved correctly!")
}

// Note: We may add another function by following the same way above

func main() {
	// Create a customer
	c1 := customer{firstName: "Gorkem", lastName: "Gul", age: 24, gender: "male"}
	c1.save()

	// Create another customer
	c2 := customer{firstName: "Recep", lastName: "Kaya", age: 24, gender: "male"}
	c2.save()
}
