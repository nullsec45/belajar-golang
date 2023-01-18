package main

import "fmt"

func getFullName() (string, string, string) {
	return "Rama", "Fajar", "Fadhillah"
}

func main() {
	//firstName, lastName, middleName := getFullName()
	//fmt.Println(firstName, lastName, middleName)

	//tanda _, untuk menghiraukan data selanjutnya tidak akan dipakai
	firstName, _, _ := getFullName()
	fmt.Println(firstName)
}
