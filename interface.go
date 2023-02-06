package main

import "fmt"

type HasName interface {
	GetName() string
	GetGender() string
}

type Person struct {
	Name   string
	Gender string
}

type Animal struct {
	Name   string
	Gender string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func AnimalDesc(hasName HasName) {
	fmt.Println("Nama Hewan", hasName.GetName())
	fmt.Println("Jenis Kelamin", hasName.GetGender())
}
func (person Person) GetName() string {
	return person.Name
}

func (person Person) GetGender() string {
	return person.Gender
}

func (animal Animal) GetName() string {
	return animal.Name
}
func (animal Animal) GetGender() string {
	return animal.Gender
}

func main() {
	var fajar Person
	fajar.Name = "Fajar"
	fajar.Gender = "Laki-laki"
	var entong Person
	entong.Name = "Entong"
	entong.Gender = "Laki-laki"

	SayHello(fajar)
	SayHello(entong)

	var cat Animal
	cat.Name = "Cat"
	cat.Gender = "Jantan"

	dog := Animal{
		Name:   "Blacki",
		Gender: "Jantan",
	}
	AnimalDesc(cat)
	AnimalDesc(dog)
}
