package sort

import (
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
	"time"
)

const ARR_LENGTH = 4000
const ARR_MAX = 100000

func TestInsertSort(t *testing.T) {
	Convey("BubbleSort", t, func() {
		Convey("ASC", func() {
			var arr []int
			for i := 0; i < ARR_LENGTH; i++ {
				rand.Seed(time.Now().UnixNano())
				arr = append(arr, rand.Intn(ARR_MAX))
			}
			intArr := NewIntArr(arr)
			BubbleSort(intArr, SortAsc)
			So(intArr.CheckSort(SortAsc), ShouldBeTrue)
		})

		Convey("DESC", func() {
			var arr []int
			for i := 0; i < ARR_LENGTH; i++ {
				rand.Seed(time.Now().UnixNano())
				arr = append(arr, rand.Intn(ARR_MAX))
			}
			intArr := NewIntArr(arr)
			BubbleSort(intArr, SortDesc)
			So(intArr.CheckSort(SortDesc), ShouldBeTrue)
		})
	})

	Convey("InsertSort", t, func() {
		Convey("ASC", func() {
			var arr []int
			for i := 0; i < ARR_LENGTH; i++ {
				rand.Seed(time.Now().UnixNano())
				arr = append(arr, rand.Intn(ARR_MAX))
			}
			intArr := NewIntArr(arr)
			InsertSort(intArr, SortAsc)
			So(intArr.CheckSort(SortAsc), ShouldBeTrue)
		})

		Convey("DESC", func() {
			var arr []int
			for i := 0; i < ARR_LENGTH; i++ {
				rand.Seed(time.Now().UnixNano())
				arr = append(arr, rand.Intn(ARR_MAX))
			}
			intArr := NewIntArr(arr)
			InsertSort(intArr, SortDesc)
			So(intArr.CheckSort(SortDesc), ShouldBeTrue)
		})
	})

	// SelectionSort
	Convey("SelectionSort", t, func() {
		Convey("ASC", func() {
			var arr []int
			for i := 0; i < ARR_LENGTH; i++ {
				rand.Seed(time.Now().UnixNano())
				arr = append(arr, rand.Intn(ARR_MAX))
			}
			intArr := NewIntArr(arr)
			SelectionSort(intArr, SortAsc)
			So(intArr.CheckSort(SortAsc), ShouldBeTrue)
		})

		Convey("DESC", func() {
			var arr []int
			for i := 0; i < ARR_LENGTH; i++ {
				rand.Seed(time.Now().UnixNano())
				arr = append(arr, rand.Intn(ARR_MAX))
			}
			intArr := NewIntArr(arr)
			SelectionSort(intArr, SortDesc)
			So(intArr.CheckSort(SortDesc), ShouldBeTrue)
		})
	})

	// MergeSort
	Convey("MergeSort", t, func() {
		Convey("ASC", func() {
			var arr []int
			for i := 0; i < ARR_LENGTH; i++ {
				rand.Seed(time.Now().UnixNano())
				arr = append(arr, rand.Intn(ARR_MAX))
			}
			intArr := NewIntArr(arr)
			MergeSort(intArr, SortAsc)
			So(intArr.CheckSort(SortAsc), ShouldBeTrue)
		})

		Convey("DESC", func() {
			var arr []int
			for i := 0; i < ARR_LENGTH; i++ {
				rand.Seed(time.Now().UnixNano())
				arr = append(arr, rand.Intn(ARR_MAX))
			}
			intArr := NewIntArr(arr)
			MergeSort(intArr, SortDesc)
			So(intArr.CheckSort(SortDesc), ShouldBeTrue)
		})
	})

	Convey("QuickSort", t, func() {
		Convey("ASC", func() {
			var arr []int
			for i := 0; i < ARR_LENGTH; i++ {
				rand.Seed(time.Now().UnixNano())
				arr = append(arr, rand.Intn(ARR_MAX))
			}
			intArr := NewIntArr(arr)
			QuickSort(intArr, SortAsc)
			So(intArr.CheckSort(SortAsc), ShouldBeTrue)
		})

		Convey("DESC", func() {
			var arr []int
			for i := 0; i < ARR_LENGTH; i++ {
				rand.Seed(time.Now().UnixNano())
				arr = append(arr, rand.Intn(ARR_MAX))
			}
			intArr := NewIntArr(arr)
			QuickSort(intArr, SortDesc)
			So(intArr.CheckSort(SortDesc), ShouldBeTrue)
		})
	})
}

func getArr() []int {
	var arr1 []int
	for i := 0; i < ARR_LENGTH; i++ {
		rand.Seed(time.Now().UnixNano())
		arr1 = append(arr1, rand.Intn(ARR_MAX))
	}
	return arr1
}
func BenchmarkBubbleSort(b *testing.B) {
	arr1 := getArr()
	var arr = make([]int, ARR_LENGTH)
	b.StopTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		copy(arr, arr1)
		intArr := NewIntArr(arr)
		BubbleSort(intArr, SortDesc)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	arr1 := getArr()
	var arr = make([]int, ARR_LENGTH)
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(arr, arr1)
		intArr := NewIntArr(arr)
		InsertSort(intArr, SortDesc)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	arr1 := getArr()
	var arr = make([]int, ARR_LENGTH)
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(arr, arr1)
		intArr := NewIntArr(arr)
		SelectionSort(intArr, SortDesc)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	arr1 := getArr()
	var arr = make([]int, ARR_LENGTH)
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(arr, arr1)
		intArr := NewIntArr(arr)
		MergeSort(intArr, SortDesc)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	arr1 := getArr()
	var arr = make([]int, ARR_LENGTH)
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(arr, arr1)
		intArr := NewIntArr(arr)
		QuickSort(intArr, SortDesc)
	}
}
