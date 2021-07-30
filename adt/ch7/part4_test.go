package ch7

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{26,5,37,1,61,11,59,15,48,19}
	//arr = []int{1,5}
	QuickSort(arr)
	t.Log(arr)
}
