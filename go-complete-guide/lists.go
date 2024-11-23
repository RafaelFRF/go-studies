package main

import (
	"example.com/complete-guide/practice"
)

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	practice.Practice()
	// prices := []float64{10.99, 8.99}
	// fmt.Println(prices[1:2])
	// prices[1] = 9.99
	// fmt.Println(prices[1:2])
	// // prices[2] = 10.0 -> Generates error

	// prices = append(prices, 5.99)
	// fmt.Println(prices)
	// prices = prices[1:]
	// fmt.Println(prices)
}

// Array
// var productNames [4]string = [4]string{"A Book"}
// prices := [4]float64{10.99, 9.99, 45.99, 20.00}
// productNames[2] = "A Carpet"
// fmt.Println(prices)
// fmt.Println(prices[1])
// fmt.Println(productNames)

//// Slice
// featurePrices := prices[1:3] // included:excluded -> cold be [1:] or [:3]
// highlightedPrices := featurePrices[:1]
// fmt.Println(featurePrices)
// fmt.Println(len(featurePrices), cap(featurePrices))
// fmt.Println(highlightedPrices)
