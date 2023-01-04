package main

import "fmt"

func main() {
	var name = "Fadhillah"

	if name == "Fajar" {
		fmt.Println("Hello Fajar")
	} else if name == "Rama" {
		fmt.Println("Hello Rama")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
