package double_list_v1

// 注意双链表和循环链表是两个东西
func NewDoubleListV1() *DoubleListV1 {
	return &DoubleListV1{}
}

type DoubleListV1 struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Prev  *Node
	Next  *Node
	Value any
}

func (list *DoubleListV1) PushFront(value any) {
	node := &Node{Value: value}
	if list.Head == nil {
		list.Head = node
		list.Tail = node
		return
	}

	node.Next = list.Head
	list.Head.Prev = node
	list.Head = node

	return
}

func (list *DoubleListV1) PushBack(value any) {

}

func (list *DoubleListV1) Keys() []any {
	result := make([]any, 0)
	for node := list.Head; node != nil; node = node.Next {
		result = append(result, node.Value)
	}
	return result
}
