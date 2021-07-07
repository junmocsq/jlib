package ch5

import "testing"

func TestNewBinTree(t *testing.T) {
	tree := NewBinTree()
	arr := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	for _, v := range arr {
		tree.Add(v)
	}
}
