package main

import "fmt"

func main() {
	name := "Fadhillah"

	switch name {
	case "Rama":
		fmt.Println("Hello Rama")
	case "Fajar":
		fmt.Println("Hello Fajar")
	case "Fadhillah":
		fmt.Println("Hello Fadhillah")
	default:
		fmt.Println("Kenalan Donk")
	}
	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama Terlalu Panjang")
	//case false:
	//	fmt.Println("Nama Sudah Benar")
	//}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
