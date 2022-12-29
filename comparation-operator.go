package main

import "fmt"

func main() {
	var firstName = "Rama"
	var lastName = "Rama"

	var result bool = firstName == lastName
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
