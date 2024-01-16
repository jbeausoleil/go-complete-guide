package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	// var investmentAmount = 1000  			// integer
	//var investmentAmount float64 = 1000 		// explicitly called float64
	var investmentAmount float64 // go expects type of return
	//expectedReturnRate := 5.5           		// float64 (:= infers the type)
	var expectedReturnRate float64
	//var years float64 = 10
	var years float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future Value:", futureValue)
	fmt.Println("Future Adjusted Value:", futureRealValue)
}
