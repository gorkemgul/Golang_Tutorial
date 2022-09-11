package main

import "fmt"

// We're gonna look at how to use the defer statement with a struct
// Let's start by creating a struct
type product struct {
	productName string
	unitPrice   int
}

// Create a function to it with the struct
func (p product) saveProduct() {
	fmt.Println("Product was saved correctly!")
	defer log()
}

// Create a log function just to show how to add log with defer into our functions
func log() {
	fmt.Println("Logged!")
}

func main() {
	// Let's create a product
	p1 := product{productName: "Computer", unitPrice: 15000}
	defer p1.saveProduct()
	fmt.Println(p1)
}
