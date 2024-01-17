package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	// var investmentAmount = 1000  			// integer
	//var investmentAmount float64 = 1000 		// explicitly called float64
	var investmentAmount float64 // go expects type of return
	//expectedReturnRate := 5.5           		// float64 (:= infers the type)
	var expectedReturnRate float64
	//var years float64 = 10
	var years float64

	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	fv, frv := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	//fmt.Println("Future Value:", futureValue)
	fmt.Printf("Future Value: %.2f\n", fv)
	fmt.Printf("Future Adjusted Value: %.2f", frv)
	//fmt.Printf("Future Value: %.2f\n"+
	//	"Future Value: %.2f\n",
	//	futureValue,
	//	futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureRealValue
}
