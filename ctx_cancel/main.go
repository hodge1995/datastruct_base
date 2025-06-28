package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	for i := 0; i < 3; i++ {
		go Worker(ctx)
	}
	time.Sleep(time.Second * 2)
	cancel()
	time.Sleep(time.Second * 5)
}

func Worker(ctx context.Context) {
	time.Sleep(time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx done")
			return
		default:
			fmt.Println("worker processing")
			time.Sleep(time.Millisecond * 100)
		}
	}
}
