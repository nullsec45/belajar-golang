package main

import "fmt"

func main() {
	type NoKTP string

	var noKTPRama NoKTP = "12345678"
	fmt.Println(noKTPRama)
}
