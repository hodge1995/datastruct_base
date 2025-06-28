package main

import (
	"fmt"
	"sync"
)

func main() {
	result := []int{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(len(result))
}
