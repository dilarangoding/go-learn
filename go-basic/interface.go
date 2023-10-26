package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Orang struct {
	Nama string
}

func (orang Orang) GetName() string {
	return orang.Nama
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {

	var zoko Orang
	zoko.Nama = "Zoko"
	SayHello(zoko)

	cat := Animal{Name: "Pusy"}
	SayHello(cat)
}
