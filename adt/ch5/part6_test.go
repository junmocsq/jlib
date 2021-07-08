package ch5

import (
	"sort"
	"testing"
)

func TestNewHeap(t *testing.T) {
	maxHeap := NewHeap(maxCompare)
	minHeap := NewHeap(minCompare)

	arr := []int{99, 172, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	maxHeap.Add(arr...)
	minHeap.Add(arr...)

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	for _, v := range arr {
		if _v, ok := maxHeap.Delete(); !ok || _v != v {
			t.Errorf("最大堆添加删除失败 %d %d", _v, v)
		}
	}

	sort.Ints(arr)
	for _, v := range arr {
		if _v, ok := minHeap.Delete(); !ok || _v != v {
			t.Errorf("最小堆添加删除失败 %d %d", _v, v)
		}
	}

	maxHeap.Add(arr...)
	//t.Log(maxHeap.arr)
	maxHeap.Change(7, 11)
	//t.Log(maxHeap.arr)
	maxHeap.DeleteWithIndex(3)
	//t.Log(maxHeap.arr)
	maxHeap.Clear()
	maxHeap.Add(arr...)
	maxLinkHeap := NewLinkHeap(maxCompare)
	maxLinkHeap.Add(arr...)
	//t.Log(maxLinkHeap.Arr(),maxHeap.arr)

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	for _, v := range arr {
		if _v, ok := maxHeap.Delete(); !ok || _v != v {
			t.Errorf("最大堆添加删除失败 %d %d", _v, v)
		}
		if _v, ok := maxLinkHeap.Delete(); !ok || _v != v {
			t.Errorf("maxLinkHeap添加删除失败 %d %d", _v, v)
		}
	}
}
