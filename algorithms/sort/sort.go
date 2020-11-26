package sort

// isAsc 是否正序
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
