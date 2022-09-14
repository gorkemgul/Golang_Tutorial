package main

import (
	"fmt"
	"os"
)

// We've tested our function and it was working correctly but what if it found the document but there was another type of error. (e.g. maybe the type of the document wasn't true or etc.)
// Here, we're gonna focus on the type assertion

// Let's re-create our openFile function
func openFile() {
	f, err := os.Open("example.txt")
	if err != nil {
		if pathErr, ok := err.(*os.PathError); ok {
			fmt.Printf("Couldn't find the document in this path! --> %v\n", pathErr.Path)
			return
		} else {
			fmt.Println("There is another type of mistake")
			return
		}
	} else {
		fmt.Println(f.Name())
	}
}

func main() {
	openFile()
}
