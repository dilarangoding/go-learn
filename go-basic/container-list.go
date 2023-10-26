package main

import (
	"container/list"
	"fmt"
)

func main() {

	data := list.New()

	data.PushBack("John")
	data.PushBack("Bowo")
	data.PushBack("Zoko")
	data.PushFront("Ahmad")

	//for element := data.Front(); element != nil; element = element.Next() {
	//	fmt.Println(element.Value)
	//}

	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
