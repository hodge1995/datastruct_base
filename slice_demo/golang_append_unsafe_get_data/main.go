package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 1. 将地址字符串转换为 uintptr（实际使用时替换为你的地址）
	addr := uintptr(0x140000162b8) // 实际地址从你的场景获取

	// 2. 将 uintptr 转换为 unsafe.Pointer
	ptr := unsafe.Pointer(addr) // nolint

	// 3. 将 unsafe.Pointer 转换为 *int 指针
	intPtr := (*int)(ptr)

	// 4. 通过指针解引用获取值
	value := *intPtr

	fmt.Printf("地址 %#x 的值是: %d\n", addr, value)
}
