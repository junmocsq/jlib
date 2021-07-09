package ch5

import "testing"

func TestBinTreeFind(t *testing.T) {
	tree := NewBinTree()
	arr := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	for _, v := range arr {
		tree.InsertRecv(v)
	}
	tree.InOrder()
	tree.LevelOrderIndex()
	tree.Insert(8)
	tree.Insert(8)
	tree.Delete(8)
	tree.InOrder()
	tree.LevelOrderIndex()
	tree.Insert(7)
	tree.Insert(7)
	tree.Insert(7)
	tree.InsertRecv(7)
	tree.DeleteRecv(7)
	tree.DeleteRecv(8)
	tree.DeleteRecv(8)
	tree.InOrder()
	tree.LevelOrderIndex()
}

func BenchmarkBinTree_Insert(b *testing.B) {
	b.StopTimer()
	arr := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	tree := NewBinTree()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		tree.Clear()
		b.StartTimer()
		for _, v := range arr {
			tree.InsertIter(v)
		}
	}

}

func BenchmarkBinTree_InsertRecv(b *testing.B) {
	b.StopTimer()
	arr := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	tree := NewBinTree()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		tree.Clear()
		b.StartTimer()
		for _, v := range arr {
			tree.InsertRecv(v)
		}
	}
}
