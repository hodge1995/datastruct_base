package main

import "math/rand"

const maxLevel = 32
const pFactor = 0.25

type SkiplistNode struct {
	val     int
	forward []*SkiplistNode // 下一个列, 数值相等
}

type Skiplist struct {
	head  *SkiplistNode
	level int
}

func Constructor() Skiplist {
	return Skiplist{
		&SkiplistNode{
			-1,
			make([]*SkiplistNode, maxLevel)},
		0,
	}
}

func (Skiplist) randomLevel() int {
	level := 1
	for level < maxLevel && rand.Float64() < pFactor {
		level++
	}
	return level
}

// 这块代码实际画图走了一遍流程还是很复杂的。
//head 节点并不是第一个带有value的节点
//forward可以理解为当前节点的下一个数组的一个指针
//这个查找的方法和理论上是一样的, 不用去曲解
func (s *Skiplist) Search(target int) bool {
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 target 的元素
		for curr.forward[i] != nil && curr.forward[i].val < target {
			curr = curr.forward[i]
		}
	}
	curr = curr.forward[0]
	// 检测当前元素的值是否等于 target
	return curr != nil && curr.val == target
}

func (s *Skiplist) Add(num int) {
	update := make([]*SkiplistNode, maxLevel)
	for i := range update {
		update[i] = s.head
	}
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 num 的元素
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	lv := s.randomLevel()
	s.level = max(s.level, lv)
	newNode := &SkiplistNode{num, make([]*SkiplistNode, lv)}
	for i, node := range update[:lv] {
		// 对第 i 层的状态进行更新，将当前元素的 forward 指向新的节点
		newNode.forward[i] = node.forward[i]
		node.forward[i] = newNode
	}
}

func (s *Skiplist) Erase(num int) bool {
	update := make([]*SkiplistNode, maxLevel)
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 num 的元素
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	curr = curr.forward[0]
	// 如果值不存在则返回 false
	if curr == nil || curr.val != num {
		return false
	}
	for i := 0; i < s.level && update[i].forward[i] == curr; i++ {
		// 对第 i 层的状态进行更新，将 forward 指向被删除节点的下一跳
		update[i].forward[i] = curr.forward[i]
	}
	// 更新当前的 level
	for s.level > 1 && s.head.forward[s.level-1] == nil {
		s.level--
	}
	return true
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
