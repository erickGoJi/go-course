package user

import (
	"errors"
	"fmt"
	"time"
)

// User is a struct, it will hold the user data
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// Struct embedding(in other langauges it is called inheritance)
type Admin struct {
	email    string
	password string
	User
}

// (u user) is a receiver
func (u User) OutputUserDatails() {
	fmt.Println("First Name: ", u.firstName)
	fmt.Println("Last Name: ", u.lastName)
	fmt.Println("Birth Date: ", u.birthDate)
	fmt.Println("Created At: ", u.createdAt)
}

// NewAdmin is a constructor/creation function
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "User",
			birthDate: "01/01/2000",
			createdAt: time.Now(),
		},
	}
}

// newUser is a constructor/creation function
func New(firstName, lastName, birthDate string) (*User, error) {
	// validate the input data
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("all fields are required")
	}
	// return a pointer to the struct and nil error
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

// (u *user) is a pointer receiver, it will modify the original struct
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
