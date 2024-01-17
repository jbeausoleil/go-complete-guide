package main

import (
	"fmt"
)

// Goals
// 1. Ask for revenue, expenses and tax rate
// 2. Calculate earnings before tax (EBT) and earnings after tax (profit)
// 3. Calculate ratio (EBT / Profit)
// 4. Output EBT, profit and ratio

func main() {

	// Ask for revenue, expenses and tax rate
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	earningsBeforeTax, profit, expenseRatio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Expense Ratio: %.2f", expenseRatio)
}

func getUserInput(infoText string) float64 {
	var userInput float64

	fmt.Print(infoText)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return 0
	}

	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	// Calculate earnings before tax (EBT) and earnings after tax (profit)
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)

	// Calculate ratio (EBT / Profit)
	expenseRatio := earningsBeforeTax / profit
	return earningsBeforeTax, profit, expenseRatio
}
