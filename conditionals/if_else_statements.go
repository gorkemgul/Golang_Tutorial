package main

import "fmt"

func main() {
	// Let's say we have a bank account to save money, and one day we need to take out some money. We may see the possibilities below:
	// - If we don't have enough money to take out we'll get an error.
	// - If we have enough money then we can take out how much money we want

	// As we can see, the story is the same as before, but this time we're gonna use an else-if in case of the amount and the money on the account are equal
	// Let's create the variables
	var account float64 = 10000
	var amount_of_money float64 = 9000

	// Check if the amount is higher or not
	if amount_of_money > account {
		fmt.Println("You don't have enough money on your account!")
	} else if amount_of_money == account { // Check if they're equal
		fmt.Println("Your money is being prepared, Do not forget to take your credit-card!")
		fmt.Println("Warning you have taken out all of your money!")
	} else { // Check if the amount is lower
		fmt.Println("Your money is being prepared, Do not forget to take your credit-card!")
	}
}
