package go_lru_v1

import (
	"container/list"
)

type Cache struct {
	cap     int
	lruList *list.List
	mMap    map[any]*list.Element
}

// Entry is a single entry in the cache.
type Entry struct {
	key   any
	value any
}

func NewCache(cap int) *Cache {
	return &Cache{
		cap:     cap,
		lruList: list.New(),
		mMap:    make(map[any]*list.Element),
	}
}

func (c *Cache) Keys() []any {
	keys := make([]any, 0, c.lruList.Len())
	for e := c.lruList.Front(); e != nil; e = e.Next() {
		keys = append(keys, e.Value.(*Entry).key)
	}
	return keys
}

func (c *Cache) Set(key, value any) {
	if c.lruList.Len() >= c.cap {
		c.RemoveOldest()
	}
	c.add(key, value)
}

func (c *Cache) add(key, value any) {
	ele := c.lruList.PushFront(&Entry{key: key, value: value})
	c.mMap[key] = ele
}

func (c *Cache) Get(key any) (any, bool) {
	if ele, ok := c.mMap[key]; ok {
		c.lruList.MoveToFront(ele)
		return ele.Value.(*Entry).value, true
	}
	return nil, false
}

func (c *Cache) Remove(key any) {
	if ele, ok := c.mMap[key]; ok {
		c.lruList.Remove(ele)
		delete(c.mMap, key)
	}
}

func (c *Cache) RemoveOldest() {
	if ele := c.lruList.Back(); ele != nil {
		c.lruList.Remove(ele)
		delete(c.mMap, ele.Value.(*Entry).key)
	}
}
