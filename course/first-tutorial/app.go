package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 6.5
	var investmentAmount float64
	var expectedReturnRate float64 = 5.5
	var years float64 = 10

	fmt.Println("Input an investment amount")
	fmt.Scan(&investmentAmount)

	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue, realFutureValue)
}
