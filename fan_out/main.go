package main

import (
	"fmt"
	"time"
)

func main() {
	in := make(chan int, 10)
	for i := 0; i < 3; i++ {
		go func() {
			for j := 0; j < 50; j++ {
				in <- j + i*10
			}
		}()
	}
	time.Sleep(1 * time.Second)

	out := make(chan int, 10)
	go func() {
		for v := range in {
			out <- v
		}
	}()

	go func() {
		for v := range out {
			fmt.Println(v)
		}
	}()

	time.Sleep(10 * time.Second)
}
