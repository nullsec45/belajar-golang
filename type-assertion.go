package main

import "fmt"

func random() interface{} {
	//return "Ups"
	return 100
}

func main() {
	var result interface{} = random()
	//var resultStr string = result.(string)
	//fmt.Println(resultStr)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}

}
