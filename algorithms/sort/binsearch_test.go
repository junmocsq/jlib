package sort

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 4, 5, 6, 11, 13, 15, 17, 88, 99, 100, 200, 300, 305, 1000, 2000, 2003, 9999, 12232}

	searchArr := []int{0, 1, 98, 10000, 12232, 12233}
	resArr := []int{-1, 0, -1, -1, len(arr) - 1, -1}

	for k, searchEle := range searchArr {
		if res := binarySearchBase(arr, searchEle); res != resArr[k] {
			t.Errorf("search ele:%d want:%d actual:%d", searchEle, resArr[k], res)
		}
	}

	for k, searchEle := range searchArr {
		if res := binarySearchLoopBase(arr, searchEle); res != resArr[k] {
			t.Errorf("search ele:%d want:%d actual:%d", searchEle, resArr[k], res)
		}
	}

	arr = []int{1, 1, 3, 4, 4, 4, 5, 5, 5, 5, 6, 11, 13, 15, 17, 88, 99}

	searchArr = []int{0, 1, 4, 5, 99, 12233}
	resArr = []int{-1, 0, 3, 6, len(arr) - 1, -1}
	for k, searchEle := range searchArr {
		if res := binarySearchLoopBaseHead(arr, searchEle); res != resArr[k] {
			t.Errorf("search ele:%d want:%d actual:%d", searchEle, resArr[k], res)
		}
	}

	searchArr = []int{0, 1, 4, 5, 99, 12233}
	resArr = []int{-1, 1, 5, 9, len(arr) - 1, -1}
	for k, searchEle := range searchArr {
		if res := binarySearchLoopBaseTail(arr, searchEle); res != resArr[k] {
			t.Errorf("search ele:%d want:%d actual:%d", searchEle, resArr[k], res)
		}
	}

	arr = []int{1, 1, 3, 4, 4, 4, 5, 5, 5, 5, 7, 7, 7, 7, 11, 13, 15, 17, 88, 99}
	searchArr = []int{-1, 1, 4, 5, 6, 99, 12233}
	resArr = []int{0, 0, 3, 6, 10, len(arr) - 1, -1}
	// 查找第一个大于等于给定值的元素
	for k, searchEle := range searchArr {
		if res := binarySearchLoopBaseHeadGTE(arr, searchEle); res != resArr[k] {
			t.Errorf("search ele:%d want:%d actual:%d", searchEle, resArr[k], res)
		}
	}

	arr = []int{1, 1, 3, 4, 4, 4, 5, 5, 5, 5, 7, 7, 7, 7, 11, 13, 15, 17, 88, 99}
	searchArr = []int{0, 1, 4, 5, 6, 99, 12233}
	resArr = []int{-1, 1, 5, 9, 9, len(arr) - 1, len(arr) - 1}
	// 查找最后一个小于等于给定值的元素
	for k, searchEle := range searchArr {
		if res := binarySearchLoopBaseTailLTE(arr, searchEle); res != resArr[k] {
			t.Errorf("search ele:%d want:%d actual:%d", searchEle, resArr[k], res)
		}
	}
}
