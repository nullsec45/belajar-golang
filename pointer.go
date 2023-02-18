package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Jakarta Timur", "DKI Jakarta", "Indonesia"}
	address2 := address1

	address2.City = "Jakarta Selatan"
	fmt.Println(address1)
	fmt.Println(address2)

	//	Pointer
	var address3 *Address = &address1

	*address3 = Address{"Solo", "Jawa Tengah", "Indonesia"}
	//address3.City = "Jakarta Pusat"

	var address4 *Address = new(Address)
	fmt.Println(address1)
	fmt.Println(address3)
	fmt.Println(address4)
}
