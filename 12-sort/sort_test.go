package sort

import (
	"fmt"
	"reflect"
	"testing"
)

type TESTS struct {
	arr []int
	res []int
}

var tests []TESTS

func init() {

	tests = append(tests,
		TESTS{
			[]int{1},
			[]int{1},
		},
		TESTS{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
		TESTS{
			[]int{2, 3, 5, 6, 1, 7, 4},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		TESTS{
			[]int{1, 100, 50, 40},
			[]int{1, 40, 50, 100},
		},
		TESTS{
			[]int{99, 88, 77, 66, 44, 55, 22, 33},
			[]int{22, 33, 44, 55, 66, 77, 88, 99},
		})
}
func TestMergeSort(t *testing.T) {
	for _, test := range tests {
		fmt.Println("before sort: ", test.arr)
		res := mergeSort(test.arr)
		fmt.Println("after  sort: ", res)
		if !reflect.DeepEqual(res, test.res) {
			t.Fail()
		}
	}
}

func TestQuickSort(t *testing.T) {
	for _, test := range tests {
		fmt.Println("before sort: ", test.arr)
		res := quickSort(test.arr)
		fmt.Println("after  sort: ", res)
		if !reflect.DeepEqual(res, test.res) {
			t.Fail()
		}
	}
}

func TestQuick2Sort(t *testing.T) {
	for _, test := range tests {
		fmt.Println("before sort: ", test.arr)
		res := quick2Sort(test.arr)
		fmt.Println("after  sort: ", res)
		if !reflect.DeepEqual(res, test.res) {
			t.Fail()
		}
	}
}
