package main

import "fmt"

func main() {
	age := 44 // regular variable

	agePointer := &age // pointer variable

	fmt.Println("Age: ", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println("Adult years: ", age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
