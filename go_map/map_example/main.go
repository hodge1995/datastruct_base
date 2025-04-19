package main

import (
	"fmt"

	"datastruct_base/go_map/map_v1"
)

func main() {
	m := map_v1.NewMapV1(10)

	// Set
	m.Put("hello", "world")
	m.Put("eat", "food")
	m.Put("ping", "pong")

	// Keys
	fmt.Println("keys: ", m.Keys()) // keys: [{eat food} {hello world} {ping pong}]

	// Get
	fmt.Println("key: hello, value: ", m.Get("hello")) // key: none, value:  <nil>
	fmt.Println("key: eat, value: ", m.Get("eat"))     // key: eat, value:  food
	fmt.Println("key: none, value: ", m.Get("none"))   // key: none, value:  <nil>
}
