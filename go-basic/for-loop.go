package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {

		fmt.Println("Perulangan ke ", counter)

	}

	slice := []string{"John", "Ahmad", "Lennon", "Budo", "Agus"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println("-------------------")

	for _, value := range slice {
		//fmt.Println("Index", i, "=", value)
		fmt.Println(value)
	}

	fmt.Println("------------")

	person := make(map[string]string)
	person["name"] = "John"
	person["title"] = "musician"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
