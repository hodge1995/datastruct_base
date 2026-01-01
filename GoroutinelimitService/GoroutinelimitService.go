package GoroutinelimitService

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func LimitGoroutine() {
	ch := make(chan int, 3)
	startGoNums := runtime.NumGoroutine()

	for i := 0; i < 10; i++ {
		ch <- i
		wg.Add(1)
		go PrintNumber(&wg, i, ch, startGoNums)
	}

	wg.Wait()

}

func PrintNumber(wg *sync.WaitGroup, i int, ch chan int, startGoNums int) {
	defer wg.Done()
	time.Sleep(time.Second * 3)
	fmt.Println("NumGo: ", runtime.NumGoroutine()-startGoNums)
	<-ch
}
