package DicttreeService

import (
	"fmt"
	"testing"
)

func TestLimitGoroutine(t *testing.T) {

	arrList := []string{"how", "hi", "her", "hello", "so", "see"}
	tree := Constructor()
	//添加跟节点
	for _, value := range arrList {
		tree.Insert(value)
	}

	fmt.Println("hi | true", tree.Search("hi"))
	fmt.Println("ho | false", tree.Search("ho"))

	//tree := createTree()
	//fmt.Println(tree)
	//flag := tree.findWord("her")
	//fmt.Println(flag)
	//flag = tree.findWord("hello")
	//fmt.Println(flag)
	//flag = tree.findWord("sllo")
	//fmt.Println(flag)
}

