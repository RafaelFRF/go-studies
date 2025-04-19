package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	// doubled := doubleNumbers(&numbers)
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	// transformedFn1 := getTransformedFunction(&numbers)
	// transformedNumbers := transformNumbers(&numbers, transformedFn1)
	transformedNumbers := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	transformedFn2 := getTransformedFunction(&moreNumbers)
	moredTransformedNumberss := transformNumbers(&numbers, transformedFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moredTransformedNumberss)
}

// func doubleNumbers(numbers *[]int) []int {
// 	dNumbers := []int{}
// 	for _, val := range *numbers {
// 		dNumbers = append(dNumbers, double(val))
// 	}
// 	return dNumbers
// }

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func getTransformedFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
