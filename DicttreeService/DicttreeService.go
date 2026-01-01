package DicttreeService

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

// boolean startsWith(String prefix) 如果之前已经插入的字符串word 的前缀之一为 prefix ，返回 true ；否则，返回 false
func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

//func createTree() TrieTree {
//	arrList := []string{"how", "hi", "her", "hello", "so", "see"}
//	myTree := TrieTree{}
//	//添加跟节点
//	myTree.root = &TrieNode{}
//	for _, value := range arrList {
//		myTree.addWord(value)
//	}
//	return myTree
//}
