package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile string = "balance.txt"

func readFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to read file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return value, nil
}

func writeValueToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func main() {
	var accBalance, err = readFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		panic("Cant continue sorry")
	}

	fmt.Println("Welcome to go bank")

	for {
		fmt.Println("What do you wish to do")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("your choice: ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			var accBalance, err = readFloatFromFile(accountBalanceFile)
			if err != nil {
				fmt.Println("error reading: ", err)
				return
			}
			fmt.Printf("Balance is %2.f", accBalance)
		case 2:
			fmt.Print("your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			accBalance += depositAmount
			writeValueToFile(accBalance, accountBalanceFile)
			var accBalance, err = readFloatFromFile(accountBalanceFile)
			if err != nil {
				fmt.Println("error reading: ", err)
				return
			}
			fmt.Printf("Balance is updated: %2.f", accBalance)

		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("invalid amount")
				continue
			}

			if withdrawalAmount > accBalance {
				fmt.Println("invalid amount")
				continue
			}

			accBalance -= withdrawalAmount
			writeValueToFile(accBalance, accountBalanceFile)
			var accBalance, err = readFloatFromFile(accountBalanceFile)
			if err != nil {
				fmt.Println("error reading: ", err)
				return
			}
			fmt.Printf("Balance is updated %2.f", accBalance)

		default:

			fmt.Println("Goodbye")
			fmt.Println("Thanks for coming")

			return
		}
	}
}