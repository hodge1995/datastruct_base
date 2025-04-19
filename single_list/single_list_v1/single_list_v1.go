package single_list_v1

type SingleList struct {
	Head *Node
}

type Node struct {
	Next  *Node
	Value any
}

func NewSingleList() *SingleList {
	return &SingleList{}
}

func (list *SingleList) PushFront(value any) {
	node := &Node{
		Value: value,
	}
	node.Next = list.Head
	list.Head = node
}

func (list *SingleList) PushBack(value any) {
	node := &Node{
		Value: value,
	}
	if list.Head == nil {
		list.Head = node
		return
	}

	currentNode := list.Head

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = node
}

func (list *SingleList) Keys() []any {
	keys := make([]any, 0)
	for node := list.Head; node != nil; node = node.Next {
		keys = append(keys, node.Value)
	}
	return keys
}
