package main

// Interfaces are named collections of method signatures.

import (
	"fmt"
	"math"
)

// Create an interface
type shape interface {
	area() float64
}

// Define a couple of structs to work with
type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// Implement the methods of our shape interface
func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Create a function to calculate the area of a shape
func test(s shape) {
	fmt.Printf("The area of the shape is %v\n", s.area())
}

func main() {

	// Let's test it on a rectangle
	r1 := rectangle{width: 10, height: 6}
	test(r1)

	// Let's test it on a circle
	c1 := circle{radius: 10}
	test(c1)
}
