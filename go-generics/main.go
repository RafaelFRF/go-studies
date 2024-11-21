package main

func man() {
	result := add(1, 2)
}

func add(a, b interface{}) interface{} {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)
	if aIsInt && bIsInt {
		return aInt + bInt
	}
	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)
	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}
}
