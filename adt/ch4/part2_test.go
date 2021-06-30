package ch4

import "testing"

func TestNewList(t *testing.T) {
	list := NewList()
	arr := []string{"bat", "cat", "mat", "sat", "vat"}
	list.Add(arr...)
	if !list.Equal(arr) {
		t.Errorf("链表添加失败")
	}
	list.Add(arr...)
	if 2 != list.DeleteWithNode("vat") {
		t.Errorf("链表删除失败")
	}
	if 0 != list.DeleteWithNode("bats") {
		t.Errorf("链表删除失败")
	}

	list.Clear()
	list.Add(arr...)
	list.DeleteOddNode()
	if !list.Equal([]string{"cat", "sat"}) {
		t.Errorf("链表删除奇数个节点失败失败")
	}

	list.Clear()
	list.Add(arr...)
	list.DeleteEvenNode()
	if !list.Equal([]string{"bat", "mat", "vat"}) {
		t.Errorf("链表删除偶数个节点失败失败")
	}

	list.Clear()
	list2 := NewList()
	a1 := []string{"d", "e", "f", "g", "h", "j", "l", "m", "z"}
	a2 := []string{"a", "b", "c", "d", "i", "k", "x"}
	list.Add(a1...)
	list2.Add(a2...)
	list.MergeSortList(list2)
	list.Print()
	list2.Print()

	list.Clear()
	list2.Clear()
	a1 = []string{"x1", "x2", "x3", "x4"}
	a2 = []string{"y1", "y2", "y3"}
	list.Add(a1...)
	list2.Add(a2...)
	list.MergeCross(list2)
	list.Print()
	list.Move("x1", 3)
	list.Print()
}
