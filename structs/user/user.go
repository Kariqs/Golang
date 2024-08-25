package user

import (
	"errors"
	"fmt"
	"time"
)

// This is a struct
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// Struct embedding- Add the features of one atruct to another
type Admin struct {
	email    string
	password string
	User
}

// Constructor fo Admin
func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthdate: "_____",
			createdAt: time.Now(),
		},
	}
}

// A method attached to the struct to print the details in the struct
func (user User) PrintUserDetails() {
	fmt.Println(user.firstName, user.lastName, user.birthdate)
}

// Clear user name - struct mutation by pointing to its address
func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

// Creation function or a constructor
func New(firstname, lastname, birthdate string) (*User, error) {
	//A constructor can also be used for validation
	if firstname == "" || lastname == "" || birthdate == "" {
		return nil, errors.New("Kindly fill in all the inputs.")
	}
	return &User{
		firstName: firstname,
		lastName:  lastname,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
