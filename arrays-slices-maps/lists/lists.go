package main

import "fmt"

// Dynamic Arrays
func main() {
	prices := []float64{10.99, 9.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99)
	fmt.Println(prices)
	// remove first element
	prices = prices[1:]
	fmt.Println(prices)

	discountedPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountedPrices...)
	fmt.Println(prices)
}

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
