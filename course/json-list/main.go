package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"json-list/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving note failed")
		return
	}

	fmt.Println("Saving successfull")
}

func getNoteData() (string, string) {
	title := getUserInput("Title: ")
	content := getUserInput("Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}