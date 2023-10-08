package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Ahmad"
	names[1] = "Joko"
	names[2] = "Ahmad"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{89, 20, 30}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string

	fmt.Println(len(lagi))

	var nilai = [...]int{30, 40, 50, 20}

	fmt.Println(nilai)
	fmt.Println(nilai[0])
	fmt.Println(nilai[1])
	fmt.Println(nilai[2])
	fmt.Println(nilai[3])

}
