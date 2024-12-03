package main

import (
	"fmt"
	"time"
)

type user struct {
	first_name string
	last_name  string
	birth      string
	createdAt  time.Time
}

func getUserInput(infoText string) (string, error) {
	var userInput string
	fmt.Print(infoText)

	fmt.Scan(&userInput)

	return userInput, nil
}

func output_user_details(u *user) {
	fmt.Printf("greetings %s %s, you joined at %s, your birth is on: %s ", (*u).first_name, (*u).last_name, (*u).createdAt, (*u).birth)
}

func (u user) output_reciever() {
	fmt.Printf("greetings %s %s, you joined at %s, your birth is on: %s ", u.first_name, u.last_name, u.createdAt, u.birth)
}

func main() {

	first_name_val, err := getUserInput("Enter first name: ")
	if err != nil {
		println("WRONG")
	}
	last_name_val, err := getUserInput("enter last name: ")

	if err != nil {
		println("wrong again")
	}

	birth_val, err := getUserInput("enter birth: ")

	if err != nil {
		println("wrong again")
	}

	var appUser = user{
		first_name: first_name_val,
		last_name:  last_name_val,
		birth:      birth_val,
		createdAt:  time.Now(),
	}

	output_user_details(&appUser)
	appUser.output_reciever()
}
