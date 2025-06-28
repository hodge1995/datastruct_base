package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "你好"
	fmt.Println("字节长度 (len):", len(s))                                  // 输出 27 (bytes)
	fmt.Println("字符数量 (RuneCountInString):", utf8.RuneCountInString(s)) // 输出 15 (runes/characters)
	// 也可以通过转换为 []rune 来获取字符数量
	fmt.Println("字符数量 ([]rune):", len([]rune(s))) // 输出 15
}
