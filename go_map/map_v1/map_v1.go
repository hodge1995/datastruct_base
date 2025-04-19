package map_v1

import (
	"container/list"
	"crypto/md5"
	"fmt"
)

type MapValue struct {
	Key   any
	value any
}

func NewMapV1(cap int) *MapV1 {
	var arr []*list.List
	for i := 0; i < cap; i++ {
		l := list.New()
		ll := l.Init()
		arr = append(arr, ll)
	}

	return &MapV1{
		cap:   cap,
		array: arr,
	}
}

type MapV1 struct {
	cap   int
	array []*list.List
}

func (m *MapV1) hash(key any) int {
	node, err := md5.New().Write([]byte(fmt.Sprintf("%v", key)))
	if err != nil {
		panic(err)
	}
	return node
}

func (m *MapV1) Put(key, value any) {
	node := m.hash(key)
	if m.array[node] == nil {
		m.array[node] = list.New()
	}

	m.array[node].PushFront(MapValue{key, value})
	return
}

func (m *MapV1) Get(key any) any {
	node := m.hash(key)
	if m.array[node] == nil {
		return nil
	}
	l := m.array[node]
	for e := l.Front(); e != nil; e = e.Next() {
		mapNode := e.Value.(MapValue)
		if mapNode.Key == key {
			return mapNode.value
		}
	}
	return nil
}

func (m *MapV1) Keys() []any {
	arr := make([]any, 0)
	length := len(m.array)
	for i := 0; i < length; i++ {
		l := m.array[i]
		for e := l.Front(); e != nil; e = e.Next() {
			if e.Value != nil {
				arr = append(arr, e.Value)
			}
		}
	}
	return arr
}
