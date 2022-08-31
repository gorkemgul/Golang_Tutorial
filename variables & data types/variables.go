package main

import "fmt"

func main() {
	// Define a string called text
	var text string = "Hello World!"
	fmt.Println(text)

	// Define an integer called number
	var number int = 10
	fmt.Println(number)

	// Define a float called decimal_number
	var decimal_number float32 = 10.2
	fmt.Println(decimal_number)

	// When defining a variable we don't have to follow the ways showed above, we could use another way. (:=)
	variable := 25
	fmt.Println(variable)
	// Note: When using this way to create a variable if we enter a number, a string or a float and etc. We don't need to write it's type it realizes the type by itself.
	// To check the type of our variable we use "%T" & printf, Let's see how to check the type of our variable.
	fmt.Printf("Type of our variable is: %T\n", variable)

	// Let's define another variable with the same way above
	decimal_number2 := 25.5
	// Print the value of decimal_number2 out
	fmt.Printf("Type of decimal_number2 is %T\n", decimal_number2)

	// Define a boolean variable
	var status bool = true
	fmt.Println(status)

	// Let's create two string and compare them
	var name1 string = "Gorkem"
	var name2 string = "Yasemin"

	// Define a boolean variable called status2
	var status2 bool

	// Compare them (Check if they're equal or not)
	status2 = name1 == name2
	fmt.Println(status2)

	// Create another variable called status 3 and check if they're not equal using "!="
	var status3 bool
	status3 = name1 != name2
	fmt.Println(status3)

	// To do a not operation we use "!" when printing it
	fmt.Println(!status3) // It's gonna print false
}
