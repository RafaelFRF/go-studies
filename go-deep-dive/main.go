package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	// sum := sumup(numbers)
	sum := sumup(1, 10, 15, 40, -5)
	anotherSum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// func sumup(numbers []int) int {
func sumup(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, val := range numbers {
		sum += val // sum = sum + val
	}

	return sum
}
