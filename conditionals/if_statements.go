package main

import "fmt"

func main() {
	// Let's say we have a bank account to save money, and one day we need to take out some money. We may see the possibilities below:
	// - If we don't have enough money to take out we'll get an error.
	// - If we have enough money then we can take out how much money we want

	// Create the variables
	var account float64 = 1000
	var amount_of_money float64 = 900

	// Check if the amounf is higher
	if amount_of_money > account {
		fmt.Println("You don't have enough money on your account!")
	}

	// Check if the amount is lower
	if amount_of_money <= account {
		fmt.Println("Your money is being prepared, Do not forget to take your credit-card!")
		account = account - amount_of_money
	}

	// To check how much money we have left
	fmt.Printf("You have %v dollars left", account)
}
