package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {

	users := []User{
		{"Joko", 10},
		{"Ahmad", 23},
		{"Susi", 32},
		{"Bowo", 31},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)

}
