package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser user

	appUser := user{
		userFirstName,
		userLastName,
		userBirthDate,
		time.Now(),
	}

	// ... do something awesome with that gathered data!

	// fmt.Println(firstName, lastName, birthdate)
	outputUserDetails(&appUser)
}

func outputUserDetails(u *user) {
	// ....
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
