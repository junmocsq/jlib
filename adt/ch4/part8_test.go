package ch4

import "testing"

func TestNewDoubleLink(t *testing.T) {
	d := NewDoubleLink()
	arr := []string{"a1", "a2", "a3", "a4", "a5", "a6", "a7"}
	arr2 := []string{"a7", "a6", "a5", "a4", "a3", "a2", "a1"}
	d.AddFronts(arr...)
	d.Print()
	if !d.Equal(arr2) {
		t.Error("AddFronts failed")
	}
	d.Clear()
	d.AddRears(arr...)
	if !d.Equal(arr) {
		t.Error("AddRears failed")
	}
	d.Delete("a1")
	if !d.Equal(arr[1:]) {
		t.Error("Delete failed")
	}
	d.Delete("a7")
	if !d.Equal(arr[1:6]) {
		t.Error("Delete failed")
	}
	d.AddFronts(arr...)
}

func TestNewDoubleLoopLink(t *testing.T) {
	d := NewDoubleLoopLink()
	arr := []string{"a1", "a2", "a3", "a4", "a5", "a6", "a7"}
	arr2 := []string{"a7", "a6", "a5", "a4", "a3", "a2", "a1"}
	d.AddFronts(arr...)
	if !d.Equal(arr2) {
		t.Error("AddFronts failed")
	}
	d.Clear()
	d.AddRears(arr...)
	if !d.Equal(arr) {
		t.Error("AddRears failed")
	}
	d.Delete("a1")
	if !d.Equal(arr[1:]) {
		t.Error("Delete failed")
	}
	d.Delete("a7")
	if !d.Equal(arr[1:6]) {
		t.Error("Delete failed")
	}
	d.AddFronts(arr...)
}
