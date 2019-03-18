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
