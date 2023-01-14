package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}
	slice := []string{"Rama", "Fajar", "Fadhillah", "Budi", "Joko"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Rama"
	person["title"] = "Web Developer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
