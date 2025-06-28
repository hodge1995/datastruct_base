package main

import (
	"fmt"
	"time"
)

func main() {
	in := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
	}()

	out := make(chan int, 10)
	for i := 0; i < 3; i++ {
		go func() {
			for v := range in {
				out <- v
				time.Sleep(1 * time.Second)
			}
		}()
	}

	go func() {
		for v := range out {
			fmt.Println(v)
		}
	}()
	time.Sleep(5 * time.Second)
}
