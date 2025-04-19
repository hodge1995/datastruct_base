package main

import (
	"fmt"

	"datastruct_base/single_list/single_list_v1"
)

func main() {
	singleList := single_list_v1.NewSingleList()
	fmt.Println(singleList.Keys())
	singleList.PushFront(4)
	fmt.Println(singleList.Keys())
	singleList.PushFront(3)
	fmt.Println(singleList.Keys())
	singleList.PushFront(2)
	fmt.Println(singleList.Keys())
	singleList.PushFront(1)
	fmt.Println(singleList.Keys())

	singleList.PushBack(0)
	fmt.Println(singleList.Keys())
	singleList.PushBack(-1)
	fmt.Println(singleList.Keys())
	singleList.PushBack(-2)
	fmt.Println(singleList.Keys())
}
