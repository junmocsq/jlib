package algorithms

type fRecursion struct {
	arr map[int]int
}

func NewF() *fRecursion {
	return &fRecursion{
		arr: make(map[int]int),
	}
}

func (f *fRecursion) f(key int) int {
	var ff func(int) int
	ff = func(key int) int {
		if key == 1 {
			return 1
		}
		if key == 2 {
			return 2
		}
		return ff(key-1) + ff(key-2)
	}
	return ff(key)
}

func (f *fRecursion) fUnRepeat(key int) int {
	var ff func(int) int
	ff = func(key int) int {
		if key == 1 {
			return 1
		}
		if key == 2 {
			return 2
		}

		if r, ok := f.getRes(key); ok {
			return r
		}
		sum := ff(key-1) + ff(key-2)
		f.arr[key] = sum
		return sum
	}
	return ff(key)
}

func (f *fRecursion) getRes(key int) (int, bool) {
	r, ok := f.arr[key]
	return r, ok
}
