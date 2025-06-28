package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 100; i < 110; i++ {
		wg.Add(1)
		go Worker(i, &wg)
	}
	wg.Wait()
}

func Worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	println("Worker", id, "is working")
}
