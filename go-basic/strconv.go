package main

import (
	"fmt"
	"strconv"
)

func main() {

	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("10000000000", 10, 64)

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(10000000, 10)
	fmt.Println(value)

	valInt, _ := strconv.Atoi("2000000")

	fmt.Println(valInt)

}
