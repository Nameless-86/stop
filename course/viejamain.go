package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Println("Input revenue, expenses and taxRate")
	fmt.Scan(&revenue)
	fmt.Scan(&expenses)
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax - (earningsBeforeTax * taxRate)

	ratio := earningsBeforeTax / profit

	fmt.Printf("Earnings before tax %.03f", earningsBeforeTax)
	fmt.Printf("Profit: %f", profit)
	fmt.Printf("ratio: %f", ratio)

}
