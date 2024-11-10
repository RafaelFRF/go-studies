package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// Receiver argument | Function attached to the struct
func (u *user) outputUserDetails() {
	// Correct way to read
	// fmt.Println((*u).firstName, (*u).lastName, (*u).birthdate)
	// Short way Go permit/recomended to use
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthdate string) (*user, error) {
	// Validation step
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println("------------------------")
		fmt.Println(err)
		fmt.Println("------------------------")
		return
	}

	// appUser := user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// }
	// Make sure it's all ordened to make this short way
	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now(),
	// }

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
