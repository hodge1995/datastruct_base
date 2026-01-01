package OddEvenService

import (
	"fmt"
	"time"
)

func PrintOddEven() {

	ch := make(chan int, 1)
	go odd(ch)
	go even(ch)

	time.Sleep(time.Second * 10)
}

func odd(ch chan int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			ch <- 1
			time.Sleep(time.Second)
		}
	}
}

func even(ch chan int) {
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			<-ch
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}
}
