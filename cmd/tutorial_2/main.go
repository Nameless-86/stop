package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int32 = 32767
	intNum += 1
	fmt.Println(intNum)

	var floatNum float32 = 123456.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(len("ñ"))
	fmt.Println(utf8.RuneCountInString("ñ"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 rune
	fmt.Println(intNum3)

	myVar := "text"
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "constant"
	fmt.Println(myConst)

	const pi float32 = 3.1415
}
