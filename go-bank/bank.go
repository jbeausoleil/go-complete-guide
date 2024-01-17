package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")

	//for i := 0; i < 2; i++ {
	for {
		fmt.Println("What do you want to do?")

		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your account balance is:", accountBalance)
		} else if choice == 2 {
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
		} else if choice == 3 {
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
		} else {
			fmt.Println("Goodbye.")
			//return // will not allow code outside of loop to be executed
			break // exit program - will allow outside of loop to be executed
		}

		//fmt.Println("Your choice: ", choice)
	}
	fmt.Println("Thanks for choosing our bank")
}
