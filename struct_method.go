package main

import "fmt"

type CustomerMethod struct {
	Name, Address string
	Age           int
}

// cara 1
//func sayHelloMethod(customer CustomerMethod, name string) {
//	fmt.Println("Hello", name, "My Name is", customer.Name)
//}

// cara 2 (lebih mudah)
func (customer CustomerMethod) sayHelloMethod(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}
func main() {
	var person CustomerMethod
	person.Name = "Fajar"
	person.Address = "Jakarta"
	person.Age = 19
	//cara 1
	//sayHelloMethod(person, "Entong")
	//cara 2 (lebih mudah)
	person.sayHelloMethod("Entong")
}
