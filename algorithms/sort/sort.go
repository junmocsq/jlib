package sort

// 冒泡排序
// 稳定的原地排序，不会改变相同数据的顺序
func BubbleSort(arr Sorter, sortFlag string) {
	length := arr.Length()
	for i := length - 1; i > 0; i-- {
		flag := false // 是否存在交换
		for m := 0; m < i; m++ {
			if sortFlag == SortAsc && arr.Compare(m, m+1) == 1 ||
				sortFlag == SortDesc && arr.Compare(m, m+1) == -1 {
				arr.Swap(m, m+1)
				flag = true // 有交换
			}
		}
		if !flag { // 如果不存在交换，说明已经排好序了
			break
		}
	}
}

// isAsc 是否正序
// 稳定的原地排序，不会改变相同数据的顺序
func InsertSort(arr Sorter, sortFlag string) {
	length := arr.Length()
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if sortFlag == SortAsc && arr.Compare(j+1, j) == -1 ||
				sortFlag == SortDesc && arr.Compare(j+1, j) == 1 {
				arr.Swap(j+1, j)
			} else {
				break
			}
		}
	}
}

// 选择排序
// 不稳定的原地排序 例如，[5，8，5，2，9]，第一次排序5和2换就乱序了
func SelectionSort(arr Sorter, sortFlag string) {
	length := arr.Length()
	for m := 0; m < length-1; m++ {
		std := m
		for n := m + 1; n < length; n++ {
			if sortFlag == SortAsc && arr.Compare(std, n) == 1 ||
				sortFlag == SortDesc && arr.Compare(std, n) == -1 {
				std = n
			}
		}
		if std != m {
			arr.Swap(std, m)
		}
	}
}

// 希尔排序 Shellsort https://zh.wikipedia.org/wiki/%E5%B8%8C%E5%B0%94%E6%8E%92%E5%BA%8F

// 归并排序
func MergeSort(arr Sorter, sortFlag string) {
	if arr, ok := arr.(*IntArr); ok {
		mergeSort(arr.arr, sortFlag)
	}
}

func mergeSort(arr []int, sortFlag string) {
	var f func(start, end int)
	var mg func(start, mid, end int)
	f = func(start, end int) {
		if start >= end {
			return
		}
		mid := (start + end) / 2
		f(start, mid)
		f(mid+1, end)
		mg(start, mid, end)
	}
	mg = func(start, mid, end int) {
		temp := make([]int, end-start+1)
		m, n := start, mid+1
		for i := 0; i < end-start+1; i++ {
			if m > mid {
				temp[i] = arr[n]
				n++
			} else if n > end {
				temp[i] = arr[m]
				m++
			} else {
				if sortFlag == "ASC" {
					if arr[m] > arr[n] {
						temp[i] = arr[n]
						n++
					} else {
						temp[i] = arr[m]
						m++
					}
				} else {
					if arr[m] > arr[n] {
						temp[i] = arr[m]
						m++
					} else {
						temp[i] = arr[n]
						n++
					}
				}
			}
		}
		for i := 0; i < end-start+1; i++ {
			arr[start+i] = temp[i]
		}
	}
	f(0, len(arr)-1)
}

func QuickSort(arr Sorter, sortFlag string) {
	if arr, ok := arr.(*IntArr); ok {
		quickSort(arr.arr, sortFlag)
	}
}

func quickSort(arr []int, sortFlag string) {

	var f func(start, end int)

	f = func(start, end int) {
		if start >= end {
			return
		}
		pivot := (start + end) / 2
		arr[end], arr[pivot] = arr[end], arr[pivot]

		// 我们通过游标 i 把 A[p...r-1]分成两部分。A[p...i-1]的元素都是小于 pivot 的，我们暂且叫它“已处理区间”，A[i...r-1]是“未处理区间”。
		// 我们每次都从未处理的区间 A[i...r-1]中取一个元素 A[j]，与 pivot 对比，如果小于 pivot，则将其加入到已处理区间的尾部，也就是 A[i]的位置。
		i, j := start, start
		for j < end {
			if sortFlag == "ASC" {
				if arr[j] < arr[end] {
					arr[i], arr[j] = arr[j], arr[i]
					i++
				}
			} else {
				if arr[j] > arr[end] {
					arr[i], arr[j] = arr[j], arr[i]
					i++
				}
			}
			j++
		}
		arr[i], arr[end] = arr[end], arr[i]
		f(start, i-1)
		f(i+1, end)
	}
	f(0, len(arr)-1)

}

// 桶排序 按范围分桶，桶里面的数据采用快排等其他方法排序，
func BucketSort(arr Sorter, sortFlag string) {
	if arr, ok := arr.(*IntArr); ok {
		bucketSort(arr.arr, sortFlag)
	}
}
func bucketSort(arr []int, sortFlag string) {
	if len(arr) < 1 {
		return
	}
	min := arr[0]
	max := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	bucketMax := max/100 + 1
	bucketIndex := func(val int) int {
		return val / 100
	}
	tempArr := make([][]int, bucketMax+1)
	for _, v := range arr {
		index := bucketIndex(v)
		tempArr[index] = append(tempArr[index], v)
	}

	if sortFlag == SortAsc {
		var i int
		for _, v := range tempArr {
			quickSort(v, sortFlag)
			for _, _v := range v {
				arr[i] = _v
				i++
			}
		}
	} else {
		var i int
		for index := bucketMax; index >= 0; index-- {
			v := tempArr[index]
			quickSort(v, sortFlag)
			for _, _v := range v {
				arr[i] = _v
				i++
			}
		}
	}
}

// 计数排序
func CountingSort(arr Sorter, sortFlag string) {
	if arr, ok := arr.(*IntArr); ok {
		countingSort(arr.arr, sortFlag)
	}
}

func countingSort(arr []int, sortFlag string) {
	if len(arr) < 1 {
		return
	}
	min := arr[0]
	max := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	countingArr := make([]int, max-min+1)
	for _, v := range arr {
		index := v - min
		countingArr[index]++
	}

	for i := 1; i < max-min+1; i++ {
		countingArr[i] += countingArr[i-1]
	}
	if sortFlag == SortAsc {
		var i int
		var pre int
		for k, v := range countingArr {
			val := k + min
			for m := 0; m < v-pre; m++ {
				arr[i] = val
				i++
			}
			pre = v
		}
	} else {
		var i = len(arr) - 1
		var pre int
		for k, v := range countingArr {
			val := k + min
			for m := 0; m < v-pre; m++ {
				arr[i] = val
				i--
			}
			pre = v
		}
	}
}

// TODO 基数排序
// 基数排序对要排序的数据是有要求的，需要可以分割出独立的“位”来比较，而且位之间有递进的关系，如果 a 数据的高位比 b 数据大，那剩下的低位就不用比较了。
// 除此之外，每一位的数据范围不能太大，要可以用线性排序算法来排序，否则，基数排序的时间复杂度就无法做到 O(n) 了。
func radixSort(arr []int, sortFlag string) {

}
