package DoubleList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	doubleList := Construct()
	doubleList.TailAppend(1)
	doubleList.TailAppend(2)
	doubleList.TailAppend(3)
	assert.Equal(t, []int{1, 2, 3}, doubleList.RangePrint())
	assert.Equal(t, 3, doubleList.Len())
}

func TestInsert(t *testing.T) {
	doubleList := Construct()
	doubleList.HeadInsert(1)
	doubleList.HeadInsert(2)
	doubleList.HeadInsert(3)
	assert.Equal(t, []int{3, 2, 1}, doubleList.RangePrint())
}
