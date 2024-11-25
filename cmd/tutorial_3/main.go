package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello Friend"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err int = intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("Result of integer division is %v", result)
	} else {
		fmt.Printf("Result of integer division is %v with remainder %v", result, remainder)
	}
	fmt.Printf("Result of integer division %v with remainder %v ", result, remainder)
}

// switch remainder {
// case 0:
// 	fmt.Printf("Exact Division")
// case 1,2:
// 	fmt.Printf("Division was close")
// default:
// 	fmt.Printf("Division was not close")

// }

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int) {
	var err error
	if denominator == 0 {
		err = errors.New("Cant divide by Zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
