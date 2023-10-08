package main

import "fmt"

func main() {

	//var person map[string]string  = map[string]string{}
	person := map[string]string{
		"name":    "John",
		"address": "Liverpool",
	}

	person["title"] = "Musician"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	fmt.Println(len(person))

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Google"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
