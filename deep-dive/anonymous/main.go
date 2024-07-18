package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3}
	// Anonymous function
	transformedNumbers := transformNumbers(numbers, func(n int) int {
		return n * 2
	})

	double := createTransformer(2)
	triple := createTransformer(3)

	doubled := transformNumbers(numbers, double)
	tripled := transformNumbers(numbers, triple)

	fmt.Println(transformedNumbers)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

// Recieves a slice of numbers and a function that takes an int and returns an int
func transformNumbers(numbers []int, f func(int) int) []int {
	var transformedNumbers []int
	for _, n := range numbers {
		transformedNumbers = append(transformedNumbers, f(n))
	}
	return transformedNumbers
}

// Higher order function, using closure
func createTransformer(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}
