package ch5

import (
	"reflect"
	"sort"
	"testing"
)

func TestNewBinTreadTree(t *testing.T) {
	tree := NewBinTreadTree()
	arr := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}

	arr1 := make([]int, len(arr))
	copy(arr1, arr)
	sort.Ints(arr1)

	tree.Add(arr...)
	if !reflect.DeepEqual(tree.ThreadInOrder(), arr1) {
		t.Error("thread binary tree failed!")
	}
	tree.insertRight(tree.root.left, 999)
	tree.insertLeft(tree.root.left, 888)
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 888, 8, 999, 9, 10, 11, 12, 13, 14, 15}
	if !reflect.DeepEqual(tree.ThreadInOrder(), arr2) {
		t.Error("thread binary tree failed!")
	}

}
