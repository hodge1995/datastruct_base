package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	arr := []int{11, 12, 13, 14, 15}

	for _, v := range arr {
		if v == 14 {
			break
		}
		fmt.Printf("Value: %d\n", v)
	}
	for v := range slices.Values(arr) {
		if v == 14 {
			break
		}
		fmt.Printf("Value: %d\n", v)
	}
	for k, v := range slices.All(arr) {
		fmt.Printf("Index: %d, Value: %d\n", k, v)
	}
	mmap := make(map[int]int)
	mmap[1] = 100
	mmap[2] = 200
	for k, v := range maps.All(mmap) {
		fmt.Printf("Map Key: %d, Value: %d\n", k, v)
	}
	for k := range maps.Keys(mmap) {
		fmt.Printf("Map Key: %d\n", k)
	}
	for v := range maps.Values(mmap) {
		fmt.Printf("Value: %d\n", v)
	}

}
