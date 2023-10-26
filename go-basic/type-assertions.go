package main

import "fmt"

func Random() interface{} {
	return true
}

func main() {

	var result interface{} = Random()

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is String")
	case int:
		fmt.Println("Value", value, "is Integer")
	default:
		fmt.Println("Unknown type")
	}

}
