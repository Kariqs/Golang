package main

import (
	"bank.com/bank-demo/fileops"
	"fmt"
)

func deposit(balance float64) {
	var depositedAmount float64
	fmt.Print("Enter amount to deposit:")
	fmt.Scan(&depositedAmount)
	if depositedAmount < 50 {
		fmt.Println("Unable to deposit KSH.", depositedAmount, ".The lowest amount you can deposit is KSH.50")
		return
	} else {
		newAccountBalance := balance + depositedAmount
		fileops.WriteFloatToFile(newAccountBalance, accountBalanceFile)
		fmt.Println("Deposit of KSH.", depositedAmount, "was successful.Your new account balance is KSH.", newAccountBalance)
	}
}

func withdraw(balance float64) {
	var withdrawedAmount float64
	fmt.Print("Enter amount to withdraw:")
	fmt.Scan(&withdrawedAmount)
	if withdrawedAmount < 50 {
		fmt.Println("Unable to withdraw KSH.", withdrawedAmount, ".The lowest amount you can withdraw is KSH.50")
		return
	} else if balance < withdrawedAmount {
		fmt.Println("You do not have enough balance to withdraw KSH.", withdrawedAmount, ".Your balance is KSHS.", balance)
		return
	} else if balance > withdrawedAmount {
		balance -= withdrawedAmount
		fileops.WriteFloatToFile(balance, accountBalanceFile)
		fmt.Println("Withdrawal of KSH.", withdrawedAmount, "was successful.Your new account balance is KSH.", balance)
	}
}
