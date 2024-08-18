package main

import (
	"fmt"
	"math"
)

func main() {
	profitCalculator()
}

func calculate() {
	const inflationRate float64 = 2.5
	name := "Bena"
	var investmentAmount float64
	var expactedReturnRate float64
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expactedReturnRate)
	fmt.Print("Number of Years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expactedReturnRate/100, years)
	var futureRealvalue = futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Hello", name, ",your future value is:", futureValue, "\nThe real value after inflation is", futureRealvalue)
}

func profitCalculator() {
	revenue := userInput("Revenue: ")
	expenses := userInput("Expenses: ")
	taxRates := userInput("Tax Rates: ")

	var earningsBeforeTax, earningsAfterTax, ratio = earnings(revenue, expenses, taxRates)

	fmt.Println("Earnings Before Tax (EBT):", earningsBeforeTax)
	fmt.Println("Earnings After Tax (EAT):", earningsAfterTax)
	fmt.Printf("Ratio of EBT to EAT:%.2f", ratio)
}

func userInput(text string) float64 {
	var inputText float64
	fmt.Print(text)
	fmt.Scan(&inputText)
	return inputText
}

func earnings(revenue, expenses, taxRates float64) (float64, float64, float64) {
	var ebt float64 = revenue - expenses
	var eat float64 = ebt * (1 - (taxRates / 100))
	var ratio float64 = ebt / eat
	return ebt, eat, ratio
}
