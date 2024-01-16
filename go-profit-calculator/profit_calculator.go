package main

import (
	"fmt"
	"math"
)

// Goals
// 1. Ask for revenue, expenses and tax rate
// 2. Calculate earnings before tax (EBT) and earnings after tax (profit)
// 3. Calculate ratio (EBT / Profit)
// 4. Output EBT, profit and ratio

func main() {
	// Declare data containers
	//revenue := 2000.0
	var revenue float64
	//expenses := 1000.0
	var expenses float64
	//taxRate := .24
	var taxRate float64

	// Ask for revenue, expenses and tax rate
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	// Calculate earnings before tax (EBT) and earnings after tax (profit)
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)

	// Calculate ratio (EBT / Profit)
	expenseRatio := earningsBeforeTax / profit

	fmt.Print("Earnings before tax: ")
	fmt.Println(earningsBeforeTax)

	fmt.Print("Profit: ")
	fmt.Println(profit)

	fmt.Print("Expense ratio: ")
	fmt.Println(math.Round(expenseRatio*100) / 100) // rounds to nearest two decimals
}
