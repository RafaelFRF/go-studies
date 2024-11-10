package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

// Receiver argument | Function attached to the struct
func (u *User) OutputUserDetails() {
	// Correct way to read
	// fmt.Println((*u).firstName, (*u).lastName, (*u).birthdate)
	// Short way Go permit/recomended to use
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// Receiver argument | Function attached to the struct
func (a *Admin) OutputAdminDetails() {
	// Correct way to read
	// fmt.Println((*u).firstName, (*u).lastName, (*u).birthdate)
	// Short way Go permit/recomended to use
	fmt.Println(a.email, a.password, a.firstName, a.lastName, a.birthdate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {
	// Validation step
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) (*Admin, error) {
	// Validation step
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}, nil
}
