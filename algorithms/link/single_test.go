package link

import (
	"testing"
)

func TestSingle(t *testing.T) {
	s := NewSingle()
	testLinker(s, t)

}

func TestCircular(t *testing.T) {
	s := NewCircular()
	testLinker(s, t)

}

func TestDouble(t *testing.T) {
	s := NewDouble()
	testLinker(s, t)

}

func TestDoubleCircular(t *testing.T) {
	s := NewDoubleCircular()
	testLinker(s, t)

}

func testLinker(s Linker, t *testing.T) {
	arr := []interface{}{
		"csq", "lmm", "zxf", "junmo", "lxq",
	}
	s.Add("lmm", "junmo", "lxq")
	s.Print()
	s.InsertByIndex(0, "csq")
	s.InsertByIndex(2, "zxf")
	for k, v := range arr {
		if !Equal(s.ValueOf(k), v) {
			t.Error("添加失败")
		}
		if s.Find(v) != k {
			t.Error("查找失败")
		}
	}
	s.Print()
	s.Add("lmm", "lmm", "lmm")
	index := s.Find("lmm")
	nextValue := s.ValueOf(index + 1)
	if s.Del("lmm") {
		if s.Find(nextValue) != index {
			t.Error("删除失败")
		}
	}
	t.Log("删除全部", s.DelAll("lmm"))

	if s.Find("lmm") != -1 {
		t.Error("删除全部失败")
	}
	s.Clear()
	s.Add("lmm", "lmm", "lmm", "lmm")
	s.Print()
	s.DelAll("lmm")
}
