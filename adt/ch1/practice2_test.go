package ch1

import (
	"math/rand"
	"testing"
	"time"
)

func TestP2Horner(t *testing.T) {
	var length = 10
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(10)
	}
	x := 10
	if P2HornerLoop(arr, x) != P2Horner_1(arr, x) {
		t.Errorf("Horner 迭代计算 failed!")
	}
	//t.Log(P2HornerLoop(arr, x), P2HornerRecursion(arr, x))
	if P2HornerRecursion(arr, x) != P2Horner_1(arr, x) {
		t.Errorf("Horner 递归计算 failed!")
	}
}

func TestP3(t *testing.T) {
	if len(P3(0)) != 0 {
		t.Errorf("真值组合失败")
	}
	if len(P3(4)) != 16 {
		t.Errorf("真值组合失败")
	}
}

func TestP6(t *testing.T) {
	res := make(map[int]bool)
	res[1] = true
	res[2] = true
	res[3] = true
	res[4] = true
	res[9] = false
	for k, v := range res {
		if P6(k) != v {
			t.Errorf("n:%d failed!", k)
		}
	}
}
func TestP7(t *testing.T) {
	if P7(10) < 0 {
		t.Errorf("n:%d failed!", 30)
	}
}

func TestP11Hanoi(t *testing.T) {
	n := 10
	_, num := P11Hanoi(n)
	if num != (1<<n)-1 {
		t.Error("汉诺塔移动次数计算失败")
	}
}

func TestP12Powerset(t *testing.T) {
	arr := []string{"a", "b", "c"}
	if len(P12Powerset(arr)) != 8 {
		t.Error("元素子集获取失败")
	}
}
func BenchmarkP2Horner(b *testing.B) {
	b.StopTimer()
	var length = 100
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(10)
	}
	x := 10
	b.StartTimer()
	b.Run("P2HornerLoop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			P2HornerLoop(arr, x)
		}
	})
	b.Run("P2HornerRecursion", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			P2HornerRecursion(arr, x)
		}
	})
	b.Run("P2Horner_1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			P2Horner_1(arr, x)
		}
	})
}

func BenchmarkP3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		P3(4)
	}
}
