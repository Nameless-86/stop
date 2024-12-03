package main

import "fmt"

func main() {
	fmt.Println("Hello Friend")
	var palabra string = "wasup"
	for i := 0; i < len(palabra); i++ {
		fmt.Printf("Char %c\n", palabra[i])
	}
}
