package TopK

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test2(t *testing.T) {
	h := &IntHeap{16, 4, 8, 70, 2, 36, 22, 5, 12}

	fmt.Println("\nHeap:")
	heap.Init(h)

	fmt.Printf("最大值: %d\n", (*h)[0])

	for i := 1000; i < 1010; i++ {
		if heap.Pop(h).(int) < i {
			heap.Push(h, i)
		}

	}
	//for(Pop)依次输出最小值,则相当于执行了HeapSort
	fmt.Println("\nHeap sort:")
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}

func Test(t *testing.T) {

	filepath := "./test.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)

	h := &IntHeap{1, 2, 3, 4, 5}
	heap.Init(h) // 将数组切片进行堆化

	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		number, err := strconv.Atoi(line)
		if err != nil {
			break
		}
		value := heap.Pop(h).(int)
		if value < number {
			heap.Push(h, number)
		} else {
			heap.Push(h, value)
		}
		//fmt.Println(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}
	fmt.Println("\nHeap sort:")
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
