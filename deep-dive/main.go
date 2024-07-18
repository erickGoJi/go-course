package main

import "fmt"

// Using Variadic Functions
func main() {
	numbers := []int{1, 10, 15}
	// The sumup function can take a variable number of arguments.
	sum := sumup(1, 10, 15, 40, -15)
	// The ellipsis (...) is used to unpack the numbers slice into individual arguments.
	anotherSum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// Variadic functions are functions that can take a variable number of arguments.
// The syntax for a variadic function is to use an ellipsis (...) before the type of the last parameter.
// The variadic parameter must be the last parameter in the function signature.
func sumup(startingValue int, numbers ...int) int {
	sum := startingValue
	for _, number := range numbers {
		sum += number
	}
	return sum
}
