package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	arr := NewArr(8)
	arr.Insert(7)
	arr.Print()
	arr.Insert(9)
	arr.Print()
	arr.Insert(2)
	arr.Print()
	arr.Insert(3)
	arr.Print()
	arr.Insert(1)
	arr.Print()
	arr.Insert(1)
	arr.Print()
	arr.Insert(1)
	arr.Print()
	arr.Insert(1)
	arr.Print()
	err := arr.Insert(1)
	fmt.Println(err)
	arr.Print()
	arr.Delete(3)
	arr.Print()
	arr.Delete(9)
	arr.Print()
	arr.Delete(1)
	arr.Print()
	arr.Delete(1)
	arr.Print()
}

func TestMergeSortedArray(t *testing.T) {
	a := []int{1, 3, 5, 7, 9}
	b := []int{2, 4, 6, 8, 10}
	m := MergeSortedArray(a, b)
	fmt.Printf("merged %v and %v to %v\n", a, b, m)

	a = []int{1, 2, 3, 4, 5}
	b = []int{6, 7, 8, 9, 10, 11, 12, 13}
	m = MergeSortedArray(a, b)
	fmt.Printf("merged %v and %v to %v\n", a, b, m)
}
