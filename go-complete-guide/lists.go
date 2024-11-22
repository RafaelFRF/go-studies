package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
	productNames[2] = "A Carpet"
	fmt.Println(prices[1])
	fmt.Println(productNames)
}
