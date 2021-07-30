package ch7

// 快速排序

func QuickSort(arr []int) {
	var f func(start, end int)

	f = func(start, end int) {
		if start>=end{
			return
		}
		pivot := arr[start]
		i := start+1
		j := end
		for i <= j {
			for arr[i] < pivot  {
				i++
			}
			for arr[j] > pivot  {
				j--
			}
			if i<j{
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		arr[start], arr[j] = arr[j], arr[start]
		f(start, j-1)
		f(j+1, end)
	}
	f(0, len(arr)-1)
}