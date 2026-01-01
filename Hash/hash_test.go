package Hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHashTable(t *testing.T) {
	mmap := NewHash(10)

	mmap.Set("hello", "world")
	assert.Equal(t, "world", mmap.Get("hello").value)

	mmap.Set("hello", "update")
	assert.Equal(t, "update", mmap.Get("hello").value)
}

func TestLoadFactor(t *testing.T) {
	mmap := NewHash(10)

	mmap.Set("hello1", "world")
	mmap.Set("hello2", "update")
	mmap.Set("hello3", "update")
	mmap.Set("hello4", "update")
	mmap.Set("hello5", "update")
	mmap.Set("hello6", "update")
	mmap.Set("hello7", "update")
	mmap.Set("hello8", "update")
	mmap.Set("hello9", "update")
	mmap.Set("hello10", "update")
	time.Sleep(time.Second * 2)
	mmap.Set("hello11", "update")
	mmap.Set("hello12", "update")
	mmap.Set("hello13", "update")
	mmap.Set("hello14", "update")
	mmap.Set("hello15", "update")
	mmap.Set("hello16", "update")
	mmap.Set("hello17", "update")

	time.Sleep(time.Second * 5)
}
