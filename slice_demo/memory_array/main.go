package main

import (
	"fmt"
)

func myAppend2(sl []int, elems ...int) {
	fmt.Printf("%p\n", &sl)
	sl = append(sl, elems...)
	sl[0] = 10000
	fmt.Printf("%p\n", &sl)
}

func main() {
	sl := make([]int, 4, 8)
	sl[0] = 1
	sl[1] = 2
	sl[2] = 3

	//var sl = []int{1, 2, 3}
	fmt.Printf("%p\n", &sl)
	fmt.Println("in slice:", sl)
	myAppend2(sl, 4, 5, 6)
	fmt.Printf("%p\n", &sl)
	fmt.Println("out slice:", sl)
}
