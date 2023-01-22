package main

import "fmt"

// Tidak perlu memanggil nama varibalenya lagi ketika return, karena
// sudah deklarasi pada saat inisiasi functionnya
func getFullName2() (firstName, middleName, lastName string) {
	firstName = "Rama"
	middleName = "Fajar"
	lastName = "Fadhillah"

	return
}

func main() {
	firstName, middleName, lastName := getFullName2()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
