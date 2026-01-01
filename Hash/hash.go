package Hash

import (
	"fmt"
	"log"
)

type Node struct {
	key      interface{} // key
	value    interface{} // value
	nextNode *Node       // 指向下一个节点
}

// 加载因子 = count / length; 加载因子越大，哈希表性能越差
type maps struct {
	count  uint    // 个数
	length uint    // 长度
	data   []*Node // Node 数组
}

type Hash interface {
	Set(key interface{}, value interface{})
	Get(key interface{}) *Node
	ListAll() []Node
}

func NewHash(i int) *maps {
	log.Println("初始化map,个数: ", i)
	// 若申请长度小于等于0，则默认给10
	if i <= 0 {
		i = 10
	}

	// 申请map
	var newMap maps
	newMap.count = 0                           // 当前hash数
	newMap.length = uint(i)                    // 申请的map长度
	newMap.data = make([]*Node, newMap.length) // 申请的内存
	return &newMap
}

func (mp *maps) hashCode(key interface{}) (sum int) {
	switch key.(type) {
	case int:
		sum = key.(int)
	case string:
		for _, v := range []byte(key.(string)) {
			sum += int(v)
		}
	}
	return
}

// 哈希函数计算槽位
func (mp *maps) hashFunc(key interface{}) int {
	return mp.hashCode(key) % int(mp.length)
}

func (mp *maps) saveData(node *Node, _id int) {
	if mp.data[_id] == nil {
		mp.data[_id] = node
	} else {
		if mp.data[_id].nextNode == nil {
			mp.data[_id].nextNode = node
		} else {
			tempNode := mp.data[_id].nextNode
			node.nextNode = tempNode
			mp.data[_id].nextNode = node
		}
	}

	mp.count++
	loadFactor := float64(mp.count) / float64(mp.length)
	fmt.Println("start:", mp.count, mp.length, loadFactor, node.key)
	mp.loadFactor(loadFactor)
}

// TestLoadFactor
func (mp *maps) loadFactor(loadFactor float64) {

	if loadFactor >= 1 {
		fmt.Println("---------------rehash--------------------")
		// 方案一: 禁止写入，只允许get, 复制hash的所有元素然后开辟新的空间并写入
		nodes := mp.ListAll()
		mp.length = mp.length * 2
		mp.count = 0
		mp.data = make([]*Node, mp.length)
		for _, node := range nodes {
			mp.Set(node.key, node.value)
		}
	}
}

func (mp *maps) ListAll() []Node {
	var nodes []Node

	for _, v := range mp.data {
		if nil != v {
			for v.nextNode != nil {
				nodes = append(nodes, *v)
				v = v.nextNode
			}
			nodes = append(nodes, *v)
		}
	}

	return nodes
}

func (mp *maps) Get(key interface{}) *Node {
	// 根据hash函数获取下标ID
	_id := mp.hashFunc(key)

	for curNode := mp.data[_id]; curNode != nil; curNode = curNode.nextNode {
		if curNode.key == key {
			return curNode
		}
	}
	return nil
}

func (mp *maps) Set(key interface{}, value interface{}) {
	// update
	node := mp.Get(key)
	if node != nil {
		node.value = value
		return
	}

	var newData Node
	_id := mp.hashFunc(key)

	newData.key = key
	newData.value = value
	newData.nextNode = nil

	// 插入数据
	mp.saveData(&newData, _id)
}
