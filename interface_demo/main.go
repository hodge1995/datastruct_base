package main

import (
	"fmt"
)

// interface{} 版本的通用容器 (简化)
type Container struct{ items []interface{} }

func (c *Container) Add(item interface{})      { c.items = append(c.items, item) }
func (c *Container) Get(index int) interface{} { return c.items[index] }

func main() {
	// 使用时
	c := Container{}
	c.Add(10)      // 存入 int
	c.Add("hello") // 存入 string

	// 取出时需要类型断言
	val1 := c.Get(0).(int) // 如果存入的不是 int，这里会 panic
	_ = val1

	// 安全的方式
	if val2, ok := c.Get(1).(string); ok {
		fmt.Println("Got string:", val2)
	} else {
		fmt.Println("Item at index 1 is not a string")
	}
}
