package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// userNames := []string{}
	// userNames = append(userNames, "Max")
	// userNames = append(userNames, "Manuel")

	userNames := make([]string, 2, 5) // Empty defined size array
	// userNames[0] = "Julie"  ->  Error
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Max")
	fmt.Println(userNames)
	
	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7
	// courseRatings["node"] = 4.7
	// fmt.Println(courseRatings)
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