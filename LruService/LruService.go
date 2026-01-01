package LruService

type LRUCache struct {
	size, capacity int
	maps           map[int]*DoubleList
	head, tail     *DoubleList
}

type DoubleList struct {
	key, value int
	prev, next *DoubleList
}


func Constructor(capacity int) LRUCache {
	newNode := new(DoubleList)
	newNode.key = 0
	newNode.value = 0
	l := LRUCache{
		maps:     map[int]*DoubleList{},
		head:     &DoubleList{key: 0, value: 0},
		tail:     &DoubleList{key: 0, value: 0},
		capacity: capacity,
	}
	l.head.next = l.tail

	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.maps[key]; !ok {
		return -1
	}
	node := this.maps[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.maps[key]; !ok {
		node := &DoubleList{key: key, value: value}
		this.maps[key] = node

		// 插入头部
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.maps, removed.key)
			this.size--
		}
	} else {
		// 更新值，移动到头部
		node := this.maps[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DoubleList) {
	node.prev = this.head
	node.next = this.head.next
	// 这个顺序
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DoubleList) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DoubleList) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DoubleList {
	node := this.tail.prev
	this.removeNode(node)
	return node
}