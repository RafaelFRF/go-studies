package main

import "fmt"

func recursion() {
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(number int) int {
	// result := 1

	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }

	// return result

	if number == 0 {
		return 1
	}
	return  number * factorial(number - 1)
	// This works from end to beggining
}