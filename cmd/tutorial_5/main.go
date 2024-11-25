package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "rèsumè"
	var indexed = myString[0]
	fmt.Printf(" %v, %T\n ", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\nThe length of 'myString is %v'", len(myString))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"p", "e", "s", "u", "t", "t", "i"}
	var strBuilder strings.Builder
	// var catStr = ""
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}
