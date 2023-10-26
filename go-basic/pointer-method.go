package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Di method", man.Name)
}

func main() {

	zoko := Man{"Zoko"}
	zoko.Married()

	fmt.Println(zoko.Name)
}
