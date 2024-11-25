package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const profitHistoryFile = "history.txt"

func calc_profit(revenue float64, expenses float64, taxRate float64) float64 {
	var earningsBeforeTax float64 = revenue - expenses
	return earningsBeforeTax - (earningsBeforeTax * taxRate)

}

func writeToProfit(value float64) error {
	valueText := strconv.FormatFloat(value, 'f', -1, 64)
	err := os.WriteFile(profitHistoryFile, []byte(valueText), 0644)
	if err != nil {
		return err
	}
	return nil
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value most be positive")
	}
	return userInput, nil
}

func readProfitHistory() (float64, error) {
	data, err := os.ReadFile(profitHistoryFile)

	if err != nil {
		return 0, errors.New("failed to read file")
	}

	profitText := string(data)
	profit, err := strconv.ParseFloat(profitText, 64)

	if err != nil {
		return 0, errors.New("failed to parse value")
	}

	return profit, nil
}

func main() {

	for {

		var choice int

		fmt.Println("1 to input, 2 to read, anything else to exit")
		fmt.Print("Your choice: ")

		fmt.Scan(&choice)

		switch choice {
		case 1:

			revenue, err := getUserInput("Revenue: ")
			if err != nil {
				fmt.Println("error writing Revenue")
				return
			}
			expenses, err := getUserInput("Expenses")
			if err != nil {
				fmt.Println("error writing Expenses")
			}
			taxRate, err := getUserInput("Tax rate (percentage)")
			if err != nil {
				fmt.Println("error writing tax rate")
			}

			profit := calc_profit(revenue, expenses, taxRate)
			err = writeToProfit(profit)
			if err != nil {
				fmt.Println("error writing profit")
			} else {
				fmt.Println("added to profit", profit)
			}
		case 2:
			// read profits
			readProfitValue, err := readProfitHistory()
			if err != nil {
				fmt.Println("Error reading profit", err)
			} else {
				fmt.Println("read profit: ", readProfitValue)
			}

		default:

			fmt.Println("Bye")
			return
		}

	}
}
