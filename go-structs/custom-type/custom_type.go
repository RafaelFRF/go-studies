package customtype

import "fmt"

type str string

func (text str) log() {
	fmt.Println(text)
}

func Customtype() {
	var name str

	name = "Max"

	name.log()
}
