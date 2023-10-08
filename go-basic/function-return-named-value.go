package main

import "fmt"

func getFullname2() (firstName, middleName, lastName string) {
	firstName = "Zoko"
	middleName = "Wi"
	lastName = "Dodo"

	return
}

func main() {

	a, b, c := getFullname2()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
