package main

import (
	"fmt"
	s "strings"
)

// Here, we're gonna look at how to use basic string functions.

func main() {

	// Create a string called name
	name := "Gorkem"

	// Let's use some of the string functions by creating a couple of examples
	fmt.Println(s.Count(name, "G"))    // Check the number of the Capital G in our name variable (Note: Go is case sensitive)
	fmt.Println(s.Contains(name, "C")) // Check if there is C in our name variable
	fmt.Println(s.Index(name, "o"))    // It returns the index of the first instance of substr in s, if it includes the written character. Otherwise it will return "-1"
	fmt.Println(s.ToLower(name))       // Turns all the characters into lower-case
	fmt.Println(s.ToUpper(name))       // Turns all the characters into upper-case

	// Let's create another string called surname to show other string functions
	var surname string = "Gul"
	fmt.Println(s.HasPrefix(surname, "Cab")) // It tests whether the string s begins with prefix.
	fmt.Println(s.HasSuffix(surname, "Bar")) // It tests whether the string s ends with suffix.

	// Create a string array and concatenate its elements
	characters := []string{"G", "o", "r", "k", "e", "m"}

	// We use strings.Join() function to concatenate the elements
	fmt.Println(s.Join(characters, ""))

	// Let's use a seperator this time and create another variable called newName
	newName := s.Join(characters, "*")
	fmt.Println(newName)

	// We may replace the operator with another one using strings.Replace() method
	fmt.Println(s.Replace(newName, "*", "-", -1)) // instead of using "-1" to change all seperates we could use len(newName)

	// We use strings.Split() to break a string into a list of substrings using a specified delimiter. It returns the substrings in the form of a slice
	fmt.Println(s.Split(newName, "*"))

	// Let's create another string and repeat it for a specific number of times
	helloWorld := "Hello World!"
	fmt.Println(s.Repeat(helloWorld, 5))
}
