package ch4

import "testing"

func TestListHead_Reverse(t *testing.T) {
	a := NewList()
	arr := []string{"a1", "a2", "a3", "a4", "a5", "a6", "a7"}
	a.Add(arr...)
	a.Reverse()
	a.Print()
	a2 := NewList()
	arr = []string{"b1", "b2", "b3", "b4", "b5", "b6"}
	a2.Add(arr...)
	a.Concatenate(a2)
	a.Print()
}

func TestNewLoop(t *testing.T) {
	l := NewLoop()
	arr := []string{"junmo", "csq", "zxf", "lxq", "lmm"}
	arr2 := []string{"lmm", "lxq", "zxf", "csq", "junmo"}
	l.AddFronts(arr...)
	if !l.Equal(arr2) {
		t.Error("AddFronts failed")
	}

	l.Clear()
	l.AddRears(arr...)
	if !l.Equal(arr) {
		t.Error("AddRears failed")
	}

	l2 := NewLoop()
	l2.AddRears(arr...)
	l.Concatenate(l2)
	if !l.Equal(append(arr, arr...)) {
		t.Error("Concatenate failed")
	}
	l.Clear()
	l.AddRears(arr...)
	l.Reverse()
	if !l.Equal(arr2) {
		t.Error("Reverse failed")
	}
}
