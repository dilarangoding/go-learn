package main

import "fmt"

type Person struct {
	Name string
}

func (person Person) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", person.Name)
}

func (a Person) sayHuuu() {
	fmt.Println("HUUUUU FROM", a.Name)
}

func main() {

	zoko := Person{Name: "Zoko"}
	zoko.sayHi("Bowo")
	zoko.sayHuuu()
}
