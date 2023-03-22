package helper

import "fmt"

var Application = "Belajar Golang"
var version = 1

// Jika nama variable atau function diawali dengan huruf kapital, itu artinya bisa diakses dari luar
// file atau bisa diimport.

// export function
func SayHello(name string) {
	fmt.Println("Hello ", name)

}

// function tidak diexport
func sayGoodBye(name string) {
	fmt.Println("Good Bye :", name)
}
