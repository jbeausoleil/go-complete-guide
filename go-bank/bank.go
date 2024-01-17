package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	err := os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
	if err != nil {
		return
	}
}

func main() {
	var accountBalance = getBalanceFromFile()

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")

		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is:", accountBalance)

		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount.  Deposit must be greater than 0.")
				//return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated.  Your account balance is now:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawlAmount float64
			fmt.Scan(&withdrawlAmount)

			if withdrawlAmount <= 0 {
				fmt.Println("Invalid amount.  Deposit must be greater than 0.")
				continue
			}

			if withdrawlAmount > accountBalance {
				fmt.Println("Invalid amount.  Please limit your withdraw to available funds.")
				continue
			}

			accountBalance -= withdrawlAmount
			fmt.Println("Balance updated.  Your account balance is now:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye.")
			fmt.Println("Thanks for choosing our bank")
			return // will not allow code outside of loop to be executed - must be used for switch statements
			//break // exit program - will allow outside of loop to be executed
		}
		fmt.Println("----------------------------------------")
	}
}
