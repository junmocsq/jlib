package ch7

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestSearch(t *testing.T) {
	var arr []int
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		arr = append(arr, rand.Intn(10000))
	}
	sort.Ints(arr)

	var res []int
	// 保证元素唯一
	for k, v := range arr {
		if k > 0 && v == res[len(res)-1] {
			continue
		}
		res = append(res, v)
	}
	arr = res
	for i := 0; i < len(arr); i++ {
		if SeqSearch(arr, arr[i]) != i {
			t.Error("SeqSearch failed!")
		}
		if BinSearch(arr, arr[i]) != i {
			t.Error("BinSearch failed!")
		}
		if BinSearchLoop(arr, arr[i]) != i {
			t.Error("BinSearchLoop failed!")
		}
	}

	searchNum := -99
	if SeqSearch(arr, searchNum) != -1 {
		t.Error("SeqSearch failed!")
	}
	if BinSearch(arr, searchNum) != -1 {
		t.Error("BinSearch failed!")
	}
	if BinSearchLoop(arr, searchNum) != -1 {
		t.Error("BinSearchLoop failed!")
	}
}

func TestInsertionSort(t *testing.T) {
	var arr []int
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		arr = append(arr, rand.Intn(10000))
	}
	arr1 := make([]int, 100)
	copy(arr1, arr)
	sort.Ints(arr1)
	InsertionSort(arr)
	if !reflect.DeepEqual(arr, arr1) {
		t.Error("InsertionSort failed")
	}
}

func TestQuickSort(t *testing.T) {
	var arr []int
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		arr = append(arr, rand.Intn(10000))
	}
	arr1 := make([]int, 100)
	copy(arr1, arr)
	sort.Ints(arr1)
	QuickSort(arr)
	if !reflect.DeepEqual(arr, arr1) {
		t.Error("QuickSort failed")
	}

}

func TestMergeSort(t *testing.T) {
	var arr []int
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		arr = append(arr, rand.Intn(10000))
	}
	arr1 := make([]int, 100)
	copy(arr1, arr)
	sort.Ints(arr1)
	MergeSort(arr)
	if !reflect.DeepEqual(arr, arr1) {
		t.Log(arr)
		t.Log(arr1)
		t.Error("MergeSort failed")
	}

	arr = nil
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		arr = append(arr, rand.Intn(10000))
	}
	copy(arr1, arr)
	sort.Ints(arr1)
	MergeSortLoop(arr)
	if !reflect.DeepEqual(arr, arr1) {
		t.Log(arr)
		t.Log(arr1)
		t.Error("MergeSortLoop failed")
	}

}
