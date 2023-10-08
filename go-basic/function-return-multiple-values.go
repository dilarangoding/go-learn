package main

import "fmt"

func getFullname() (string, string, int) {
	return "Ahmad", "Subandi", 30
}

func main() {

	firstName, lastName, _ := getFullname()

	fmt.Println(firstName)
	fmt.Println(lastName)

	fmt.Println(firstName, lastName)
}
