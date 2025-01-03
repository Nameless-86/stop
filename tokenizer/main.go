package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello world. How are you?"
	words := strings.Fields(text) // Word Tokenization
	fmt.Println(words)
}
