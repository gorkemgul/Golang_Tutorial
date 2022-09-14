package main

import (
	"errors"
	"fmt"
)

// Here, we're gonna look at how to use the error library. Let's say we want to write an algorithm which helps us to make a prediction to find the exact number.
// Say, When we enter a number below 1 and above 100 we want our algorithm to return an error using errors library. Let's look at how to do it.

// Create a function to make a prediction

func makePrediction(prediction int) (string, error) {
	actualNumber := 50
	if prediction < 1 || prediction > 100 {
		return "Wrong Prediction", errors.New("You need to enter a number between 1 & 100")
	} else if prediction < actualNumber {
		return "Try a higher one!", nil
	} else if prediction > actualNumber {
		return "Try a lower one!", nil
	}
	return "Predicted correctly", nil
}

func main() {
	// Let's test our function by making 3 prediction
	prediction1, error1 := makePrediction(50)
	prediction2, error2 := makePrediction(101)
	prediction3, error3 := makePrediction(-5)
	prediction4, error4 := makePrediction(51)

	// Print them out and see the results
	fmt.Printf("Prediction1: %v, Error1: %v \nPrediction2: %v, Error2: %v \nPrediction3: %v, Error3: %v \nPrediction4: %v. Error4: %v", prediction1, error1, prediction2, error2, prediction3, error3, prediction4, error4)
}
