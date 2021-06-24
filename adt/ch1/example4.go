package ch1

// 全排列
func Perm(arr []string) [][]string {
	result := make([][]string, 0)
	length := len(arr)
	var f func(ci int)
	f = func(ci int) {
		if ci == length {
			temp := make([]string, length)
			copy(temp, arr)
			result = append(result, temp)
			return
		}
		for i := ci; i < length; i++ {
			arr[ci], arr[i] = arr[i], arr[ci]
			f(ci + 1)
			arr[ci], arr[i] = arr[i], arr[ci]
		}
	}
	f(0)
	return result
}
