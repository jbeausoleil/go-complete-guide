package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1. Ask for revenue, expenses and tax rate
// 2. Calculate earnings before tax (EBT) and earnings after tax (profit)
// 3. Calculate ratio (EBT / Profit)
// 4. Output EBT, profit and ratio

func main() {

	// Ask for revenue, expenses and tax rate
	revenue, err01 := getUserInput("Revenue: ")
	expenses, err02 := getUserInput("Expenses: ")
	taxRate, err03 := getUserInput("Tax Rate: ")

	if err01 != nil || err02 != nil || err03 != nil {
		fmt.Println("Error: ", err03)
		return
	}

	earningsBeforeTax, profit, expenseRatio := calculateFinancials(revenue, expenses, taxRate)
	storeResults(earningsBeforeTax, profit, expenseRatio)

	fmt.Printf("Earnings before tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Expense Ratio: %.2f", expenseRatio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	// Validate user input and return error if negative or zero
	if userInput <= 0 {
		return 0, errors.New("value cannot be negative or zero")
	}

	return userInput, nil
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	err := os.WriteFile("results.txt", []byte(results), 0644)
	if err != nil {
		return
	}
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	// Calculate earnings before tax (EBT) and earnings after tax (profit)
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)

	// Calculate ratio (EBT / Profit)
	expenseRatio := earningsBeforeTax / profit
	return earningsBeforeTax, profit, expenseRatio
}
