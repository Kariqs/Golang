// In this package, we learn the control structures of go and their syntax.
package main

import (
	"bank.com/bank-demo/fileops"
	"fmt"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, error = fileops.GetFloatFromFile(accountBalanceFile)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println("Welcome to GoBank")
	for {
		presentOptions()

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
