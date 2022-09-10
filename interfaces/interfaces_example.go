package main

import "fmt"

// Here We're gonna do a strong example to explain the concept of interfaces. To do so, we'll create an interface to calculate how much does it cost monthly to a person if they took one credit or more.
// Define the Mortgage struct
type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

// Define the Car struct
type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

// Create the interface
type CreditCalculator interface {
	Calculate() float64
}

// Implement the methods of our interface
func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
}

// Define a function to calculate the monthly credit payment
func CalculateMonthlyPayment(credits []CreditCalculator) float64 {
	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}
	return paymentTotal
}

func main() {

	// Define the credits that a person took (In our case, let's say a person took 3 credits (2 Mortgage & 1 Car))
	firstCredit := Mortgage{rate: 10, creditPaymentTotal: 100000, address: "Palo Alto"}
	secondCredit := Mortgage{rate: 12, creditPaymentTotal: 50000, address: "Palo Alto"}
	thirdtCredit := Car{rate: 15, creditPaymentTotal: 60000, carInfo: "Range Rover"}

	// Create an array (In case, if a person got more than 1 credit)
	credits := []CreditCalculator{firstCredit, secondCredit, thirdtCredit}
	totalPaymentMonthly := CalculateMonthlyPayment(credits)

	// Let's check the totalPaymentMonthly and see the result
	fmt.Println(totalPaymentMonthly)
}
