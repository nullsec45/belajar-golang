package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		counter := 1
		fmt.Println("increment")
		counter++
	}
	increment()
	increment()
	fmt.Println(counter)
}
