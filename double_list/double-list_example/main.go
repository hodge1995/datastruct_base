package main

import (
	"fmt"

	"datastruct_base/double_list/double_list_v1"
)

func main() {
	doubleList := double_list_v1.NewDoubleListV1()

	doubleList.PushFront(1)
	fmt.Println(doubleList.Keys())
	doubleList.PushFront(2)
	fmt.Println(doubleList.Keys())
	doubleList.PushFront(3)
	fmt.Println(doubleList.Keys())
	doubleList.PushFront(4)
	fmt.Println(doubleList.Keys())
	doubleList.PushFront(5)
	fmt.Println(doubleList.Keys())
}
