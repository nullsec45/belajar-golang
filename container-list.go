package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	// PushBack tambah data ke paling belakang
	// PushFront tambah data ke paling depan
	data.PushBack("Rama")
	data.PushBack("Fajar")
	data.PushBack("Fadhillah")
	data.PushFront("Entong")

	// dari depan ke  belakang
	//for e := data.Front(); e != nil; e = e.Next() {
	//	fmt.Println(e.Value)
	//}

	// dari belakang ke depan
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
}
