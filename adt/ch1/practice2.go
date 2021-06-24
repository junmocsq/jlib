package ch1

import "fmt"

// 多项式求值
func P2HornerLoop(arr []int, x int) int {
	length := len(arr)
	result := arr[length-1]
	for n := length - 2; n >= 0; n-- {
		result = result*x + arr[n]
	}
	return result
}

func P2HornerRecursion(arr []int, x int) int {
	length := len(arr)
	var f func(n, index, res int) int
	f = func(n, index, res int) int {
		if index == 0 {
			res = arr[n-index]
		} else {
			res = res*x + arr[n-index]
		}
		if index == n {
			return res
		}
		return f(n, index+1, res)
	}
	return f(length-1, 0, 0)
}

func P2Horner_1(arr []int, x int) int {
	var f = func(exp int) int {
		result := 1
		for i := 0; i < exp; i++ {
			result *= x
		}
		return result
	}
	length := len(arr)
	result := 0
	for i := 0; i < length; i++ {
		result += arr[i] * f(i)
	}
	return result
}

func P3(n int) [][]bool {
	if n == 0 {
		return nil
	}
	var result [][]bool
	temp := make([]bool, n)
	var f func(index int)
	f = func(index int) {
		if index == n {
			t := make([]bool, n)
			copy(t, temp)
			result = append(result, t)
			return
		}
		temp[index] = false
		f(index + 1)
		temp[index] = true
		f(index + 1)
	}
	f(0)
	return result
}

func P6(n int) bool {
	if n <= 2 {
		return true
	}
	var res []int
	i := 2
	temp := n
	for {
		if temp == 1 {
			break
		}
		if temp%i == 0 {
			res = append(res, i)
			temp = temp / i
		} else {
			i++
		}
	}
	add := 0
	for _, v := range res {
		add += v
	}
	return add == n
}

func P7(n int) int {
	// 迭代
	result := 1
	for i := n; i > 0; i-- {
		result *= i
	}

	// 递归
	var f func(n int) int
	f = func(n int) int {
		if n == 1 {
			return 1
		}
		return f(n-1) * n
	}
	if result == f(n) {
		return result
	} else {
		return -1
	}
}

// 汉诺塔移动轨迹 result
// 汉诺塔移动次数num f(n) = 2*f(n-1) + 1
func P11Hanoi(n int) (result []string, num int) {
	var other = func(s, e string) string {
		arr := []string{"A", "B", "C"}
		for _, v := range arr {
			if v != s && v != e {
				return v
			}
		}
		return ""
	}
	// A B C
	var f func(n int, s, e string)
	f = func(n int, s, e string) {
		if n == 0 {
			return
		}
		// 若f(n) A=>C 则 f(n-1)先A=>B,n从A到C,f(n-1)再从B=>C
		// 若f(n) B=>C 则 f(n-1)先B=>A,n从B到C,f(n-1)再从A=>C 以此类推，总的6种转移方式
		oth := other(s, e)
		f(n-1, s, oth)
		result = append(result, fmt.Sprintf("%d:%s=>%s", n, s, e))
		num++
		f(n-1, oth, e)
	}
	f(n, "A", "C")
	return
}

// 计算n个元素的子集集合
func P12Powerset(arr []string) (result [][]string) {
	var f func(index int)
	length := len(arr)
	f = func(index int) {
		if index == length {
			var t []string
			for _, v := range arr {
				if v != "" {
					t = append(t, v)
				}
			}
			result = append(result, t)
			return
		}
		temp := arr[index]
		arr[index] = ""
		f(index + 1)
		arr[index] = temp
		f(index + 1)
	}
	f(0)
	return
}
