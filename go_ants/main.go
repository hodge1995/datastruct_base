package main

import (
	"fmt"

	"github.com/panjf2000/ants/v2"
)

func main() {
	p, err := ants.NewPool(100)
	if err != nil {
		panic(err)
	}
	defer p.Release()

	for i := 0; i < 30; i++ {
		p.Submit(func() {
			fmt.Println(i)
		})
	}

	fmt.Println(p.Free())
}
