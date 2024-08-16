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
	var revenue float64
	var expenses float64
	var taxRates float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRates)

	var earningsBeforeTax float64 = revenue - expenses
	var earningsAfterTax float64 = earningsBeforeTax * (1 - (taxRates / 100))
	var ratio float64 = earningsBeforeTax / earningsAfterTax
	fmt.Println("Earnings Before Tax (EBT):", earningsBeforeTax)
	fmt.Println("Earnings After Tax (EAT):", earningsAfterTax)
	fmt.Print("Ratio of EBT to EAT:", ratio)
}
