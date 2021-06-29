package ch3

import "testing"

func TestCreateMg(t *testing.T) {
	for i := 0; i < 100; i++ {
		mg := CreateMg(100, 100)
		//t.Log(mg)
		//t.Log(mg.arr[4][2])
		//break
		if mg.search() {
			break
		}
	}
}
