package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Go(func() {
		fmt.Println("hello")
	})

	wg.Go(func() {
		fmt.Println("world")
	})
	wg.Wait()
}

//func main() {
//	var wg sync.WaitGroup
//
//	wg.Add(1)
//	go func() {
//		fmt.Println("hello")
//		wg.Done()
//	}()
//
//	wg.Add(1)
//	go func() {
//		fmt.Println("world")
//		wg.Done()
//	}()
//
//	wg.Wait()
//}
