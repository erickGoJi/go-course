package main

import "fmt"

// Type alias
type floatMap map[string]float64

// Method for floatMap
func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)
	userNames[0] = "Julie"
	userNames[1] = "David"

	userNames = append(userNames, "John")
	userNames = append(userNames, "Jane")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["Go"] = 4.5
	courseRatings["React"] = 4.8
	courseRatings["Angular"] = 4.2

	courseRatings.output()

	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}
