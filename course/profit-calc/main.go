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

	var profit float64 = calc_profit(revenue, expenses, taxRate)
	fmt.Printf("Profit is %.3f", profit)
}

func calc_profit(revenue float64, expenses float64, taxRate float64) float64 {
	var earningsBeforeTax float64 = revenue - expenses
	return earningsBeforeTax - (earningsBeforeTax * taxRate)

}
