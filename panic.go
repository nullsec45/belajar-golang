package main

import "fmt"

func endApp() {
	//Digunakan untuk menangkap pesan dari panic, sehingga program tetap berjalan
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(false)
	//fmt.Println("Fajar")
}
