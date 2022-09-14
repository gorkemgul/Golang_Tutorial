package main

import (
	"fmt"
	"os"
)

// Here we're gonna focus on error handling by giving a simple example.
// Our example is going to about reading a file because it's kinda real world problem.

// Create a function to open the file we want to open
func openFile() {
	f, err := os.Open("example.txt")
	// Note: If it finds the document then err will be equal to nil, otherwise we're gonna get a string as an output. ("Document couldn't find!")
	if err != nil {
		fmt.Println("Document couldn't find!")
	} else {
		fmt.Println(f.Name()) // If it finds the document, it's going to get its name
	}
}

func main() {
	// Let's use our function and see the result
	openFile()
}
