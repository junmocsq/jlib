package link

import "testing"

func TestEqual(t *testing.T) {
	v1, v2 := 1, true
	Equal(v1, v2)
}

func TestSingle(t *testing.T) {
	s := NewSingle()
	testLinker(s, t)

	ss := s.(*single)

	arr := []interface{}{
		"a", "b", "c", "d", "f", "e",
	}
	s.Clear()
	s.Add(arr...)
	ss.Reverse()
	if !Equal(ss.Elements(), []interface{}{"e", "f", "d", "c", "b", "a"}) {
		t.Error("reverse failed!")
	}
	if ss.createLoop() != 2 {
		t.Error("check loop failed!")
	}

	arr = []interface{}{
		"a",
	}
	s.Clear()
	s.Add(arr...)
	ss.Reverse()
	if !Equal(ss.Elements(), []interface{}{"a"}) {
		t.Error("reverse failed!")
	}

	arr = []interface{}{
		"a", "b",
	}
	s.Clear()
	s.Add(arr...)
	ss.Reverse()
	if !Equal(ss.Elements(), []interface{}{"b", "a"}) {
		t.Error("reverse failed!")
	}

	ss.Clear()
	arrs := [][]interface{}{
		{"a", "b", "c", "d", "e", "f", "g"},
		{"a"},
	}

	for _, arr = range arrs {
		for k, b := range arr {
			ss.Clear()
			ss.Add(arr...)
			if !Equal(ss.DelDescN1(len(arr)-1-k).val, b) {
				t.Error("DelDescN1 failed!")
			}
		}

		for k, b := range arr {
			ss.Clear()
			ss.Add(arr...)
			if !Equal(ss.DelDescN2(len(arr)-1-k).val, b) {
				t.Error("DelDescN2 failed!")
			}
		}
	}

	arrs = [][]interface{}{
		{"a", "b", "c", "d", "e", "f", "g"},
		{"a", "b", "c", "d", "e", "f"},
		{"a", "b"},
		{"a"},
	}

	for _, arr = range arrs {
		ss.Clear()
		ss.Add(arr...)
		midIndex := (ss.Length() - 1) / 2
		if !Equal(ss.Middle().val, ss.ValueOf(midIndex)) {
			t.Error("Middle failed!")
		}
	}

	ss.Print()

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
	s.Add(arr...)
	if !Equal(s.Elements(), arr) {
		t.Error("batch add failed")
	}

	s.Add("zjb")
	arr = append(arr, "zjb")
	if !Equal(s.Elements(), arr) {
		t.Error("add failed")
	}
	s.Clear()
	arr = []interface{}{
		"csq", "lmm", "zxf", "junmo", "lxq", "csq",
	}
	s.Add(arr...)
	if s.Find("lxq") != 4 && s.Find("csq") != 0 {
		t.Error("find failed")
	}

	if s.Find("llll") != -1 {
		t.Error("find failed")
	}
	findAllRes := s.FindAll("csq")
	if len(findAllRes) != 2 && findAllRes[0] == 0 && findAllRes[1] == 5 {
		t.Error("findAll failed")
	}
	if s.FindAll("llll") != nil {
		t.Error("findAll failed")
	}

	s.Clear()
	arr = []interface{}{
		"csq", "lmm", "zxf", "junmo", "lxq", "zjb",
	}
	s.Add("lmm", "junmo", "lxq")

	if !s.InsertByIndex(0, "csq") || !s.InsertByIndex(2, "zxf") || !s.InsertByIndex(s.Length(), "zjb") {
		t.Error("insertByIndex failed")
	}

	if s.InsertByIndex(s.Length()+1, "zzz") {
		t.Error("insertByIndex failed")
	}

	if !Equal(s.Elements(), arr) {
		t.Error("insertByIndex out of range index")
	}

	for k := 0; k < s.Length(); k++ {
		if !Equal(s.ValueOf(k), arr[k]) {
			t.Error("valueOf failed")
		}
	}

	if s.ValueOf(s.Length()) != nil {
		t.Error("valueOf out of range index ")
	}

	s.Clear()
	s.Add("lmm", "lmm", "lmm")
	if s.Del("lmm") {
		if !Equal(s.Elements(), []interface{}{"lmm", "lmm"}) {
			t.Error("del failed")
		}
	}

	if s.DelAll("lmm") != 2 && s.Empty() {
		t.Error("delAll failed")
	}

	//s.Clear()
	arr = []interface{}{
		"csq", "csq", "lmm", "csq", "zxf", "junmo", "lxq", "zjb",
	}
	s.Add(arr...)

	if s.Del("lll") || !Equal(arr, s.Elements()) {
		t.Error("del failed")
	}
	if !s.Del("csq") || !s.Del("csq") || !Equal([]interface{}{"lmm", "csq", "zxf", "junmo", "lxq", "zjb"}, s.Elements()) {
		t.Error("del failed")
	}

	s.Clear()
	arr = []interface{}{
		"csq", "csq", "lmm", "csq", "zxf", "junmo", "lxq", "zjb",
	}
	s.Add(arr...)

	if s.DelAll("csq") != 3 || !Equal([]interface{}{"lmm", "zxf", "junmo", "lxq", "zjb"}, s.Elements()) {
		t.Error("delAll failed")
	}

	//s.Print()
	// DelByIndex

	arrs := [][]interface{}{
		{"csq", "csq", "lmm", "csq", "zxf", "junmo", "lxq", "zjb"},
		{"csq", "zjb"},
		{"csq"},
	}
	for _, arr = range arrs {
		s.Clear()
		s.Add(arr...)
		if !Equal(s.DelByIndex(0), arr[0]) {
			t.Error("delByIndex failed")
		}
		s.Clear()
		s.Add(arr...)
		last := s.Length() - 1
		if last > 0 && !Equal(s.DelByIndex(last), arr[last]) {
			t.Error("delByIndex1 failed")
		}
	}

	arrs = [][]interface{}{
		{"csq", "csq", "lmm", "csq", "zxf", "junmo", "lxq", "zjb"},
		{"csq", "zjb"},
		{"csq"},
	}
	for _, arr = range arrs {
		s.Clear()
		s.Add(arr...)
		if !Equal(s.DelHead(), arr[0]) {
			t.Error("delHead failed")
		}
		s.Clear()
		s.Add(arr...)
		last := s.Length() - 1
		if last > 0 && !Equal(s.DelTail(), arr[last]) {
			t.Error("delTail failed")
		}
	}
}
