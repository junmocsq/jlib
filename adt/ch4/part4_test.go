package ch4

import "testing"

func TestNewPoly(t *testing.T) {
	p1 := NewPoly()
	arr := [][2]int{{3, 14}, {2, 8}, {1, 7}}
	for _, v := range arr {
		p1.Add(v[0], v[1])
	}
	if p1.Print() != "3x^14 + 2x^8 + 1x^7" {
		t.Error("Add failed")
	}
	p2 := NewPoly()
	arr = [][2]int{{8, 14}, {-3, 10}, {10, 6}}
	for _, v := range arr {
		p2.Add(v[0], v[1])
	}
	if p2.Print() != "8x^14 + -3x^10 + 10x^6" {
		t.Error("Add failed")
	}
	p1.PAdd(p2)
	if p1.Print() != "11x^14 + -3x^10 + 2x^8 + 1x^7 + 10x^6" {
		t.Error("PAdd failed")
	}
	p1.Earse(8)
	if p1.Print() != "11x^14 + -3x^10 + 1x^7 + 10x^6" {
		t.Error("erase failed")
	}
}
