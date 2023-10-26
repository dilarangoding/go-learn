package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	Merried       bool
}

func main() {

	var person Customer
	person.Name = "Joko"
	person.Address = "Bandung"
	person.Age = 23

	fmt.Println(person)

	// Struct Literals

	ahmad := Customer{
		Name:    "Ahmad",
		Address: "Bandung",
		Age:     23,
	}

	fmt.Println(ahmad)

	zoko := Customer{"Zoko", "Solo", 57, false}
	fmt.Println(zoko)

}
