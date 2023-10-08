package main

import "fmt"

func main() {

	type NoKtp string
	type Married bool

	var noKtpJohn NoKtp = "2345678765433456"
	var marriedStatus Married = true

	fmt.Println(noKtpJohn)
	fmt.Println(marriedStatus)

}
