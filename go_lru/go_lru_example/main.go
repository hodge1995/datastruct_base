package main

import (
	"fmt"

	"datastruct_base/go_lru/go_lru_v1"
)

func main() {
	lruClient := go_lru_v1.NewCache(10)
	lruClient.Set("key1", "value1")
	lruClient.Get("key1")
	lruClient.Remove("key1")
	fmt.Println(lruClient.Keys())
}
