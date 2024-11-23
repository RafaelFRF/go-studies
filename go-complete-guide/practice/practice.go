package practice

import "fmt"

//  7
type Product struct {
	title string
	id    int64
	price float64
}

func Practice() {
	fmt.Println("--------Practice-------")
	// 1
	hobbies := [3]string{"Pets", "Games", "Movies"}

	// 2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// 3
	// hobbiesSliced := hobbies[0:2]
	hobbiesSliced := hobbies[:2]
	fmt.Println(hobbiesSliced)

	// 4
	// hobbiesSliced = hobbies[1:3]
	// fmt.Println(hobbiesSliced)
	hobbiesSliced = hobbiesSliced[1:3]
	fmt.Println(hobbiesSliced)

	// 5
	goalsDynamicList := []string{"Understand Go project's structure", "Learn how to build an Go API"}
	fmt.Println(goalsDynamicList)

	// 6
	goalsDynamicList[1] = "Create entire apps by myself"
	goalsDynamicList = append(goalsDynamicList, "Learn how to build an Go API")
	fmt.Println(goalsDynamicList)

	// 7
	productsList := []Product{
		{title: "dot", id: 1, price: 200.00}, {title: "djTable", id: 140, price: 1200.00},
	}
	fmt.Println(productsList)
	newProduct := Product{
		title: "laserPen", id: 28, price: 1.00,
	}
	productsList = append(productsList, newProduct)
	fmt.Println(productsList)

	fmt.Println("--------End-Practice-------")
}
