package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	fajar := Man{"Fajar"}
	fajar.Married()
	fmt.Println(fajar.Name)
}
