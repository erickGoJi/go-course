package main

import (
	"example/structs/user"
	"fmt"
)

// str is a custom type
type str string

// print is a method of the str type
func (s str) print() {
	fmt.Println(s)
}

func main() {
	firstName := getUserDate("Enter your first name: ")
	lastName := getUserDate("Enter your last name: ")
	birthDate := getUserDate("Enter your birth date (MM/DD/YYYY): ")

	appUser, err := user.New(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	// ... do something awesome with this data!

	appUser.OutputUserDatails()
	appUser.ClearUserName()
	appUser.OutputUserDatails()

	admin := user.NewAdmin("exaple@mail.com", "password")

	admin.OutputUserDatails()
	admin.ClearUserName()
	admin.OutputUserDatails()

	var s str = "Hello, World!"
	s.print()
}

func getUserDate(prompText string) string {
	fmt.Print(prompText)
	var input string
	fmt.Scanln(&input)
	return input
}
