package sort

import (
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
	"time"
)

func TestInsertSort(t *testing.T) {
	Convey("InsertSort", t, func() {
		Convey("ASC", func() {
			var arr []int
			for i := 0; i < 20; i++ {
				rand.Seed(time.Now().UnixNano())
				arr = append(arr, rand.Intn(1000))
			}
			intArr := NewIntArr(arr)
			InsertSort(intArr, SortAsc)
			So(intArr.CheckSort(SortAsc), ShouldBeTrue)
		})

		Convey("DESC", func() {
			var arr []int
			for i := 0; i < 20; i++ {
				rand.Seed(time.Now().UnixNano())
				arr = append(arr, rand.Intn(1000))
			}
			intArr := NewIntArr(arr)
			InsertSort(intArr, SortDesc)
			So(intArr.CheckSort(SortDesc), ShouldBeTrue)
		})
	})
}
