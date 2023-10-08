package main

import "fmt"

func main() {

	name := "Service"

	switch name {
	case "Home":
		fmt.Println("Page Home")
	case "About Us":
		fmt.Println("Page About Us")
	case "Contact":
		fmt.Println("Page Contact")
	default:
		fmt.Println("404 Not Found")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama Terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah benar")
	//}

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan  panjang")
	default:
		fmt.Println("Nama sudah benar")

	}

}
