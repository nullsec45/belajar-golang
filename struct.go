package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var customer Customer
	customer.Name = "Annisa Nur Almasih Witman"
	customer.Address = "Indonesia"
	customer.Age = 19

	fmt.Println(customer.Name)
	fmt.Println(customer.Address)
	fmt.Println(customer.Age)

	customer2 := Customer{
		Name:    "Rama",
		Address: "Indonesia",
		Age:     19,
	}
	fmt.Println(customer2)

	customer3 := Customer{"Fajar", "Jakarta", 19}
	fmt.Println(customer3)
}
