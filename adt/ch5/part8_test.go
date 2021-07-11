package ch5

import "testing"

func TestNewSelectTree(t *testing.T) {

	tree := NewSelectTree()
	tree.CreateSuccessTree(24)
	t.Log(tree.SortSuccessTree())

	tree = NewSelectTree()
	tree.CreateFailedTree(8)
	//t.Log(tree.SortFailedTree())

}
