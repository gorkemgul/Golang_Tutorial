package main

import "fmt"

// Here, we're gonna cover how to write custom exceptions. Again, let's say we are trying to write a program that helps us to find the actual number by letting us make predictions.
// In case, we want to add a border to our algorithm we may add it like we did before. But here we're gonna focus on writing an exception about it.

// create a struct
type borderException struct {
	parameter int
	message   string
}

// Implement the Error interface
func (b *borderException) Error() string {
	return fmt.Sprintf("The number %d %s", b.parameter, b.message)
}

// Create the function
func makePrediction(prediction int) (string, error) {
	if prediction < 1 || prediction > 100 {
		return "Wrong Prediction!", &borderException{parameter: prediction, message: "is out of range!"}
	}
	return "Predicted Correctly!", nil
}

func main() {
	// Let's test our function and see the results
	prediction1, error1 := makePrediction(101)
	fmt.Println(prediction1, error1)
}
