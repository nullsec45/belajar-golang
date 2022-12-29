package main

import "fmt"

func main() {
	//Jika tidak digunakann, golang tidak akan protes. Berbeda dengan variable ketika tidak digunakan
	//dia akan protes
	const firstName string = "Rama"
	const lastName string = "Fajar"
	const value = 1000

	//Akan error
	//firstName = "Entong"

	//Multiple Constant
	const (
		hobby    = "Ngoding"
		umur     = 19
		semester = 3
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
	fmt.Print(hobby)
	fmt.Println(umur)
	fmt.Println(semester)
}
