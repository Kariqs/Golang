package main

import (
	"errors"
	"fmt"
	"math"
	"os"
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
	revenue, revenueerr := userInput("Revenue: ")
	if revenueerr != nil {
		fmt.Println(revenueerr)
		return
	}
	expenses, expenseerr := userInput("Expenses: ")
	if expenseerr != nil {
		fmt.Println(expenseerr)
		return
	}
	taxRates, rateerr := userInput("Tax Rates: ")
	if rateerr != nil {
		fmt.Println(rateerr)
		return
	}

	var earningsBeforeTax, earningsAfterTax, ratio = earnings(revenue, expenses, taxRates)
	writeResultsToFile(earningsBeforeTax, earningsAfterTax, ratio)
	fmt.Println("Earnings Before Tax (EBT):", earningsBeforeTax)
	fmt.Println("Earnings After Tax (EAT):", earningsAfterTax)
	fmt.Printf("Ratio of EBT to EAT:%.2f", ratio)
}

func userInput(text string) (float64, error) {
	var inputText float64
	fmt.Print(text)
	fmt.Scan(&inputText)
	if inputText <= 0 {
		return 0, errors.New("invalid value, value cannot be 0 or less than 0")
	}
	return inputText, nil
}

func earnings(revenue, expenses, taxRates float64) (float64, float64, float64) {
	var ebt float64 = revenue - expenses
	var eat float64 = ebt * (1 - (taxRates / 100))
	var ratio float64 = ebt / eat
	return ebt, eat, ratio
}

func writeResultsToFile(ebt, eat, ratio float64) {
	earnings := fmt.Sprintf("Earnings Before Tax:%.2f\nEarnings After Tax:%.2f\nRatio:%.2f", ebt, eat, ratio)
	os.WriteFile("results.txt", []byte(earnings), 0644)
}
