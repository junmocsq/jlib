package ch5

import "testing"

func TestNewBinTree(t *testing.T) {
	tree := NewBinTree()
	arr := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	for _, v := range arr {
		tree.Add(v)
	}
	tree.InOrder()
	tree.IterInOrder()

	tree.PostOrder()
	tree.IterPostOrder()

	tree.PreOrder()
	tree.IterPreOrder()
	tree.IterPreOrder2()

	tree.LevelOrder()
}

func ExampleBinTree_InOrder() {
	tree := NewBinTree()
	arr := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	for _, v := range arr {
		tree.Add(v)
	}
	tree.InOrder()
	// output: 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15
}

func ExampleBinTree_PostOrder() {
	tree := NewBinTree()
	arr := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	for _, v := range arr {
		tree.Add(v)
	}
	tree.PostOrder()
	// output: 1 3 2 5 7 6 4 9 11 10 13 15 14 12 8
}
func ExampleBinTree_PreOrder() {
	tree := NewBinTree()
	arr := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	for _, v := range arr {
		tree.Add(v)
	}
	tree.PreOrder()
	// output: 8 4 2 1 3 6 5 7 12 10 9 11 14 13 15
}

func ExampleBinTree_LevelOrder() {
	tree := NewBinTree()
	arr := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	for _, v := range arr {
		tree.Add(v)
	}
	tree.LevelOrder()
	// output: 8 4 12 2 6 10 14 1 3 5 7 9 11 13 15
}
