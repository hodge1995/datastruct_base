package main

import (
	"fmt"
)

func foo(sl []int) {
	fmt.Printf("ing %v %d %d %p %p %p %p %p\n", sl, len(sl), cap(sl), &sl, &sl[0], &sl[1], &sl[2], &sl[3])
	sl = append(sl, 5)
	sl = append(sl, 6)
	sl[1] = 2000
	fmt.Printf("ing %v %d %d %p %p %p %p %p %p %p\n", sl, len(sl), cap(sl), &sl, &sl[0], &sl[1], &sl[2], &sl[3], &sl[4], &sl[5])
	sl[5] = 5000
	fmt.Printf("ing %v %d %d %p %p %p %p %p %p %p\n", sl, len(sl), cap(sl), &sl, &sl[0], &sl[1], &sl[2], &sl[3], &sl[4], &sl[5])
}

func main() {
	sl := []int{1, 2, 3}

	fmt.Printf("Before %v %d %d %p %p %p %p %p\n", sl, len(sl), cap(sl), &sl, &sl[0], &sl[1], &sl[2], &sl[3])
	foo(sl)
	fmt.Printf("Before %v %d %d %p %p %p %p %p\n", sl, len(sl), cap(sl), &sl, &sl[0], &sl[1], &sl[2], &sl[3])
}
