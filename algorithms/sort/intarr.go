package sort

var (
	SortAsc  = "ASC"
	SortDesc = "DESC"
)

type Sorter interface {
	Compare(m, n int) int // m>n 1 m=n 0 m<n -1
	Swap(m, n int)
	Length() int
	Index(int) (int, bool)
}

type IntArr struct {
	arr []int
}

var _ Sorter = &IntArr{}

func NewIntArr(arr []int) *IntArr {
	return &IntArr{
		arr: arr,
	}
}

func (i *IntArr) Compare(m, n int) int {
	if i.arr[m] > i.arr[n] {
		return 1
	} else if i.arr[m] == i.arr[n] {
		return 0
	} else {
		return -1
	}
}

func (i *IntArr) Swap(m, n int) {
	i.arr[m], i.arr[n] = i.arr[n], i.arr[m]
}

func (i *IntArr) Length() int {
	return len(i.arr)
}

func (i *IntArr) Index(index int) (int, bool) {
	if i.Length() <= index {
		return 0, false
	}
	return i.arr[index], true
}

func (i *IntArr) CheckSort(sortFlag string) bool {
	length := i.Length()
	if length <= 1 {
		return true
	}
	for m := 0; m < length-1; m++ {
		if sortFlag == SortAsc {
			if i.Compare(m, m+1) == 1 {
				return false
			}
		} else {
			if i.Compare(m, m+1) == -1 {
				return false
			}
		}
	}
	return true
}
