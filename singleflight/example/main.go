package main

import (
	"log"
	"sync"
	"time"

	"datastruct_base/singleflight/signleflight_v1"
)

var (
	wg sync.WaitGroup
	sf signleflight_v1.SingleFlight
)

func fetchResource() (any, error) {

	time.Sleep(3 * time.Second)
	return "result", nil
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(reqID int) {
			defer wg.Done()
			val, err := sf.Do("cache-key", fetchResource)
			if err != nil {
				log.Printf("Request %d error: %v", reqID, err)
				return
			}

			log.Printf("Request %d => %v", reqID, val)
		}(i)
	}

	wg.Wait()
}
