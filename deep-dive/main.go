package main

import (
	"fmt"
)

// transformFn is a type that represents a function that takes an int and returns an int
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(numbers, double)
	tripled := transformNumbers(numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
}

// transformNumbers takes a slice of numbers and a function that transforms a number and returns a slice of transformed numbers
func transformNumbers(numbers []int, f transformFn) []int {
	var dNumbers []int
	for _, n := range numbers {
		dNumbers = append(dNumbers, f(n))
	}
	return dNumbers
}

func double(n int) int {
	return n * 2
}

func triple(n int) int {
	return n * 3
}
