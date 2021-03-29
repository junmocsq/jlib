package sort

func binarySearchBase(arr []int, ele int) int {
	index := -1
	low := 0
	high := len(arr) - 1

	var search func(start, end int)
	search = func(start, end int) {
		if start > end {
			return
		}
		mid := start + ((end - start) >> 1)
		if arr[mid] == ele {
			index = mid
			return
		} else if arr[mid] > ele {
			search(start, mid-1)
		} else {
			search(mid+1, end)
		}
	}
	search(low, high)
	return index
}

func binarySearchLoopBase(arr []int, ele int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] == ele {
			return mid
		} else if arr[mid] > ele {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 查找第一个值等于给定值的元素
func binarySearchLoopBaseHead(arr []int, ele int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] == ele {
			if mid == 0 || arr[mid-1] < ele {
				return mid
			}
			high = mid - 1
		} else if arr[mid] > ele {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 查找最后一个值等于给定值的元素
func binarySearchLoopBaseTail(arr []int, ele int) int {

	length := len(arr)
	low := 0
	high := length - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] == ele {
			if mid == length-1 || arr[mid+1] > ele {
				return mid
			}
			low = mid + 1
		} else if arr[mid] > ele {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 查找第一个大于等于给定值的元素
func binarySearchLoopBaseHeadGTE(arr []int, ele int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] >= ele {
			if mid == 0 || arr[mid-1] < ele {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 查找最后一个小于等于给定值的元素
func binarySearchLoopBaseTailLTE(arr []int, ele int) int {

	length := len(arr)
	low := 0
	high := length - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > ele {
			high = mid - 1
		} else {
			if mid == length-1 || arr[mid+1] > ele {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}
