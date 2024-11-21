package main

import "fmt"

func man() {
	result := add(1, 2)
	fmt.Print(result)
}

// Add generic type for more reusable funcs
// func add[T any](a, b T) T {
// Specify the range of types that are allowed
func add[T int | float64](a, b T) T {
	// aInt, aIsInt := a.(int)
	// bInt, bIsInt := b.(int)
	// if aIsInt && bIsInt {
	// 	return aInt + bInt
	// }
	// aFloat, aIsFloat := a.(float64)
	// bFloat, bIsFloat := b.(float64)
	// if aIsFloat && bIsFloat {
	// 	return aFloat + bFloat
	// }
	return a + b
}
