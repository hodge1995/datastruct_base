package Tree

import "testing"

func TestConstruct(t *testing.T) {
	tree := Construct(3)
	tree.Add(2)
	tree.Add(1)
	tree.Add(3)
	tree.Add(5)
	tree.Add(6)
	tree.Add(8)

	tree.inorderTraversal(tree)
}

func TestConstructV2(t *testing.T) {
	tree := Construct(3)
	tree.AddV2(2)
	tree.AddV2(1)
	tree.AddV2(3)
	tree.AddV2(5)
	tree.AddV2(6)
	tree.AddV2(8)

	tree.inorderTraversal(tree)
}
