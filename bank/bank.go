// In this package, we learn the control structures of go and their syntax.
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, error = getBalanceFromFile()
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println("Welcome to GoBank")
	for {
		fmt.Println("What do you want to do today?")
		fmt.Println("1.Check Balance")
		fmt.Println("2.Make a deposit")
		fmt.Println("3.Withdraw cash")
		fmt.Println("4.Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your account balance is: KSH.%.2f\n", accountBalance)
		} else if choice == 2 {
			deposit(accountBalance)
		} else if choice == 3 {
			withdraw(accountBalance)
		} else if choice == 4 {
			fmt.Print("Thanks for choosing Go Bank")
			break
		} else {
			fmt.Println("Invalid choice. Try again")
		}

	}
}

func deposit(balance float64) {
	var depositedAmount float64
	fmt.Print("Enter amount to deposit:")
	fmt.Scan(&depositedAmount)
	if depositedAmount < 50 {
		fmt.Println("Unable to deposit KSH.", depositedAmount, ".The lowest amount you can deposit is KSH.50")
		return
	} else {
		newAccountBalance := balance + depositedAmount
		writeBalanceToFile(newAccountBalance)
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
		writeBalanceToFile(balance)
		fmt.Println("Withdrawal of KSH.", withdrawedAmount, "was successful.Your new account balance is KSH.", balance)
	}
}

func writeBalanceToFile(balance float64) {
	bal := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(bal), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, error := os.ReadFile(accountBalanceFile)
	if error != nil {
		return 0, errors.New("the balance file not found")
	}
	var textBalance = string(data)
	bal, error := strconv.ParseFloat(textBalance, 64)
	if error != nil {
		return 0, errors.New("failed to parse balance string to a float")
	}
	return bal, nil
}
