package main

import (
	"fmt"
)

func main() {
	mmap := make(map[int]struct{})

	mmap[1] = struct{}{}
	mmap[2] = struct{}{}
	mmap[5] = struct{}{}

	if mmap[1] == struct{}{} {
		fmt.Println("map struct")
	}
	if mmap[3] == struct{}{} {
		fmt.Println("map struct")
	}

	mmapBool := make(map[int]bool)

	mmapBool[1] = true
	mmapBool[2] = true
	mmapBool[5] = true

	if mmapBool[1] {
		fmt.Println("hello bool")
	}
	if mmapBool[3] {
		fmt.Println("hello bool")
	}

}
