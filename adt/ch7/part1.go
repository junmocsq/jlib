package ch7

// 顺序查找
func SeqSearch(arr []int, search int) int {
	result := -1
	for k, v := range arr {
		if v == search {
			result = k
			break
		}
	}
	return result
}

// 二分查找
func BinSearch(arr []int, search int) int {
	var binSearch func(s, e int) int
	binSearch = func(s, e int) int {
		if s > e {
			return -1
		}

		mid := (s + e) / 2
		if arr[mid] > search {
			return binSearch(s, mid-1)
		} else if arr[mid] < search {
			return binSearch(mid+1, e)
		} else {
			return mid
		}
	}
	return binSearch(0, len(arr)-1)
}

func BinSearchLoop(arr []int, search int) int {
	s, e := 0, len(arr)-1
	for s <= e {
		mid := (s + e) / 2
		if arr[mid] > search {
			e = mid - 1
		} else if arr[mid] < search {
			s = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// 类似于整理纸牌：每次取一张牌，并在取下一张牌之前，将这张牌放在适当的位置，使手中的所有纸牌按顺序排列。
func InsertionSort(arr []int) {
	var j int
	for i := 1; i < len(arr); i++ {
		temp := arr[i]

		for j = i; j > 0; j-- {
			if arr[j-1] > temp {
				arr[j] = arr[j-1]
			} else {
				break
			}
		}
		arr[j] = temp
	}
}

// 快速排序
func QuickSort(arr []int) {
	var f func(start, end int)
	f = func(start, end int) {
		if start >= end {
			return
		}
		pivot := arr[start] // 分界点 也可以用首 中 尾单个数比较，用中值来做分界点
		i := start + 1      // 小于分界点的
		j := end            // 大于分界点的
		for i <= j {
			for i <= j && arr[i] <= pivot {
				i++
			}
			for arr[j] > pivot {
				j--
			}
			if i < j {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		arr[start], arr[j] = arr[j], arr[start]
		f(start, j-1)
		f(j+1, end)
	}
	f(0, len(arr)-1)
}

// 归并排序
func MergeSort(arr []int) {
	var f func(s, e int)
	f = func(s, e int) {
		if s >= e {
			return
		} else {
			mid := (s + e) / 2
			f(s, mid)   // 左分
			f(mid+1, e) // 右分
			res := merge(arr[s:mid+1], arr[mid+1:e+1])
			for i := s; i <= e; i++ {
				arr[i] = res[i-s]
			}
			return
		}
	}
	f(0, len(arr)-1)
}

// 归并排序迭代实现
func MergeSortLoop(arr []int) {
	var mergePass func(length int)
	mergePass = func(length int) {
		var i int
		for ; i <= len(arr)-2*length; i += 2 * length {
			res := merge(arr[i:i+length], arr[i+length:i+2*length])
			for k := i; k < i+2*length; k++ {
				arr[k] = res[k-i]
			}
		}
		// 最后一个不成倍数的如果length为1倍多，则需要合并
		if i+length < len(arr) {
			res := merge(arr[i:i+length], arr[i+length:])
			for k := i; k < len(arr); k++ {
				arr[k] = res[k-i]
			}
		}
	}
	for i := 1; i < len(arr); i *= 2 {
		mergePass(i)
	}
}

func merge(arr1 []int, arr2 []int) []int {
	var res []int
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}
	res = append(res, arr1[i:]...)
	res = append(res, arr2[j:]...)
	return res
}
