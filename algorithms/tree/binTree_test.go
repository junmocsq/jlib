package tree

import "testing"

func TestNewArrayBinTree(t *testing.T) {
	arr := []string{"A", "B", "C", "D", "E", "F", "G"}
	a := NewArrayBinTree()
	for _, v := range arr {
		a.Add(v)
	}
	if a.prePrint() != "A->B->D->E->C->F->G->" {
		t.Errorf("pre want:A->B->D->E->C->F->G-> actual:%s", a.prePrint())
	}
	if a.midPrint() != "D->B->E->A->F->C->G->" {
		t.Errorf("mid want:D->B->E->A->F->C->G-> actual:%s", a.midPrint())
	}
	if a.postPrint() != "D->E->B->F->G->C->A->" {
		t.Errorf("post want:D->E->B->F->G->C->A-> actual:%s", a.postPrint())
	}

	l := NewLinkBinTree()
	arrInt := []int{4, 2, 6, 1, 3, 5, 7}
	for _, v := range arrInt {
		l.Add(v)
	}
	var s string
	s = l.prePrint()
	if s != "4->2->1->3->6->5->7->" {
		t.Errorf("pre want:4->2->1->3->6->5->7-> actual:%s", s)
	}
	s = l.midPrint()
	if s != "1->2->3->4->5->6->7->" {
		t.Errorf("pre want:1->2->3->4->5->6->7-> actual:%s", s)
	}
	s = l.postPrint()
	if s != "1->3->2->5->7->6->4->" {
		t.Errorf("pre want:1->3->2->5->7->6->4-> actual:%s", s)
	}
}
