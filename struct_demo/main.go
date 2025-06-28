package main

import (
	"fmt"
	"unsafe"
)

type DangerStruct struct {
	value struct{} // 零大小字段
	key   int64
}

func main() {

	d := DangerStruct{key: 111}

	emptyPtr := unsafe.Pointer(&d.value)

	var nextVar int64 = 999

	// 打印内存地址
	fmt.Printf("DangerStruct 地址: %p\n", &d)
	fmt.Printf("value 字段地址: %p\n", emptyPtr)
	fmt.Printf("nextVar 地址: %p\n", &nextVar)

	// 尝试通过 value 指针访问后面的内存
	// 这实际上是访问了 nextVar 的内存！
	accessedValue := *(*int64)(emptyPtr)
	fmt.Printf("通过 value 指针访问的值: 0x%X, %d\n", accessedValue, accessedValue)
}
