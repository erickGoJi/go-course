package main

import "fmt"

type Product struct {
	Title string
	ID    int
	Price float64
}

func main() {

	//  1. Create a new array (!) that contains three hobbies you have
	//     Output (print) that array in the command line.\
	fmt.Println("Exercise 1")
	hobbies := []string{"Programming", "Reading", "Gaming"}
	fmt.Println(hobbies)
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println("Exercise 2")
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])
	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	fmt.Println("Exercise 3")
	hobbiesSlice := hobbies[:2]
	fmt.Println(hobbiesSlice)
	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	fmt.Println("Exercise 4")
	hobbiesSlice = hobbies[1:]
	fmt.Println(hobbiesSlice)
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	fmt.Println("Exercise 5")
	courseGoals := []string{"Learn Go", "Build a project"}
	fmt.Println(courseGoals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	fmt.Println("Exercise 6")
	courseGoals[1] = "Learn React"
	fmt.Println(courseGoals)
	courseGoals = append(courseGoals, "Build a GRPC project")
	fmt.Println(courseGoals)
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	fmt.Println("Exercise 7")
	products := []Product{ // dynamic array
		{Title: "Book", ID: 1, Price: 10.99},
		{Title: "Carpet", ID: 2, Price: 20.99},
	}
	fmt.Println(products)
	products = append(products, Product{Title: "Table", ID: 3, Price: 30.99})
	fmt.Println(products)
}

// Dynamic Arrays
// func main() {
// 	prices := []float64{10.99, 9.99}
// 	fmt.Println(prices[0:1])
// 	prices[1] = 9.99

// 	prices = append(prices, 5.99)
// 	fmt.Println(prices)
// 	// remove first element
// 	prices = prices[1:]
// 	fmt.Println(prices)
// }

// Arrays
// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.5, 20.5, 30.5, 40.5}
// 	fmt.Println(prices)

// 	productNames[2] = "A Carpet"

// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	// Slices
// 	featurePrices := prices[1:]
// 	featurePrices[0] = 199.99
// 	highlightedPrices := prices[:1]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// }
