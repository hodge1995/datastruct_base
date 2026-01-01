package Tree

import "fmt"

type (
	TreeNode struct {
		Value int
		Left  *TreeNode
		Right *TreeNode
	}
)

func Construct(value int) *TreeNode {

	return &TreeNode{
		Value: value,
	}
}

// 单纯靠思路写出来的代码，没进行优化 version 1
func (t *TreeNode) Add(value int) {
	node := new(TreeNode)
	node.Value = value

	for t != nil {
		if t.Value > value {
			if t.Left == nil {
				t.Left = node
				break
			} else {
				t = t.Left
			}
		} else {
			if t.Right == nil {
				t.Right = node
				break
			} else {
				t = t.Right
			}
		}
	}
}

// 递归版本
func (t *TreeNode) AddV2(value int) *TreeNode {
	if t == nil {
		return Construct(value)
	}

	if value < t.Value {
		t.Left = t.Left.AddV2(value)
	} else {
		t.Right = t.Right.AddV2(value)
	}
	return t
}

func (t *TreeNode) inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	t.inorderTraversal(root.Left)
	fmt.Println(root.Value)
	t.inorderTraversal(root.Right)
}
