// when we develop executable programs, we'll use the package “main” for making the package as an executable program.
package main

// To use fmt package we need to import it
import "fmt"

// main() function is the entry point of the executable programs
func main() {
	// Print hello world using Println & Print
	// Note: The difference between them is when we use Println. It prints the str and moves to cursor to a new line. When it comes to Print, it just prints it does nothing else.
	fmt.Println("Hello World!")
	fmt.Print("Hello ")
	fmt.Print("World!")
}
