package main

import "fmt"

// Recursion is a process in which a function calls itself as a subroutine.
// This allows the function to be repeated several times, since it calls itself during its execution.
// Recursion is best applied when it makes the solution clearer and easier to understand.
// It is also useful when the problem can be broken down into smaller subproblems of the same type.
// In Go, recursion is used in many algorithms, such as the quicksort algorithm and the binary search algorithm.
func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

// Factorial is a mathematical function that multiplies a number by every number below it.
// For example, the factorial of 5 is 5 * 4 * 3 * 2 * 1 = 120.
// The factorial of 0 is 1.
func factorial(n int) int {
	// Recursive case
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
	// Base case
	// result := 1 // The factorial of 0 is 1

	// for i := 1; i <= n; i++ {
	// 	result *= i
	// }

	// return result
}
