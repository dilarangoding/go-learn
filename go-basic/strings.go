package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("John Lennon", "John"))
	fmt.Println(strings.Contains("John Lennon", "Alex"))
	fmt.Println(strings.Split("John Lennon", " "))

	fmt.Println(strings.ToLower("John Lennon"))
	fmt.Println(strings.ToUpper("John Lennon"))
	fmt.Println(strings.ToTitle("john lennon"))
	fmt.Println(strings.Trim("      John Lenonn     ", " "))
	fmt.Println(strings.ReplaceAll("JOHN JOHN OJHN", "JOHN", "Lennon"))
}
