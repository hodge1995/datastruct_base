package DoubleList

type ListNode struct {
	Value int
	Pre   *ListNode
	Next  *ListNode
}

type DoubleList struct {
	head *ListNode
	tail *ListNode
	size int
}

func Construct() DoubleList {
	list := new(DoubleList)
	list.size = 0
	list.tail = nil
	list.head = nil
	return *list
}

func (this *DoubleList) HeadInsert(value int) {
	node := new(ListNode)
	node.Value = value
	if this.head == nil {
		this.head = node
		this.tail = node
		this.size = 1
	} else {
		node.Next = this.head
		this.head.Pre = node
		this.head = node
		this.size++
	}

}

func (this *DoubleList) TailAppend(value int) {
	node := new(ListNode)
	node.Value = value
	if this.tail == nil {
		this.head = node
		this.tail = node
		this.size = 1
	} else {
		this.tail.Next = node
		node.Pre = this.tail
		this.tail = node
		this.size++
	}
}

func (this *DoubleList) RangePrint() (arr []int) {
	node := this.head
	for node != nil {
		arr = append(arr, node.Value)
		node = node.Next
	}
	return
}

func (this *DoubleList) Len() int {

	return this.size
}
