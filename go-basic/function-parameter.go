package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {

	fmt.Println("Hello", firstName, lastName)
}

func kali(a int, b int) {
	fmt.Println(a * b)
}

func main() {

	sayHelloTo("John", "Lennon")

	kali(10, 2)

}
