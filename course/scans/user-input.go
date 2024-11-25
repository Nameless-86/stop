package main

import (
	"fmt"
	"math"
)

func main() {
	var inflationRate float64
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Println("Input an investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("input years: ")
	fmt.Scan(&years)

	fmt.Print("Input expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Input Inflation Rate: ")
	fmt.Scan(&inflationRate)

	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue, realFutureValue)
}
