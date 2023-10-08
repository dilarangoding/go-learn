package main

import "fmt"

func main() {

	var nilai32 = 100000
	var nilai64 = int64(nilai32)
	var nilai8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	name := "John lennon"
	e := name[0]
	eString := string(e)

	fmt.Println(eString)
}
