package main

import (
	"belajar-golang/helper"
	"fmt"
)

func main() {
	helper.SayHello("Fajar")
	// tidak bisa import function sayGoodBye
	//helper.sayGoodbye("Entong")
	fmt.Println(helper.Application)
}
