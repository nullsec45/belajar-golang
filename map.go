package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Rama Fajar",
		"address": "Jakarta",
	}

	person["title"] = "Developer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["title"])
	fmt.Println(person["address"])

	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
}
