package main

import "fmt"

func endApp2() {
	message := recover()

	if message != nil {
		fmt.Println("Error dengan message", message)
	}

	fmt.Println("Aplikasi Selesai")

}

func runApp2(error bool) {
	defer endApp2()

	if error {
		panic("Applikasi error")
	}

	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp2(true)
	fmt.Println("Ahmad")
}
