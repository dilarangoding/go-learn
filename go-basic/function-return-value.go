package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello  " + name
	}
}

func operasi(operasi string, a int, b int) int {
	if operasi == "tambah" {
		return a + b
	} else if operasi == "bagi" {
		return a / b
	} else {
		return a - b
	}
}

func main() {

	name := "Ahmad"
	fmt.Println(getHello(name))

	fmt.Println(operasi("bagi", 40, 20))
	fmt.Println(getHello(""))
	fmt.Println(getHello("Agus"))
}
