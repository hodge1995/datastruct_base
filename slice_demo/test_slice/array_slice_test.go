package main

import (
	"testing"
)

func BenchmarkAppendWithoutCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sl []int
		for j := 0; j < 1000; j++ {
			sl = append(sl, j)
		}
	}
}

func BenchmarkAppendWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sl := make([]int, 0, 1000)
		for j := 0; j < 1000; j++ {
			sl = append(sl, j)
		}
	}
}
