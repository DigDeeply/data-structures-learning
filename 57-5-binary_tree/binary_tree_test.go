package binarytree

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	//创建
	bt := New(50)
	//添加节点
	bt.Add(8)
	bt.Add(3)
	bt.Add(9)
	bt.Add(1)
	bt.Add(7)
	bt.Add(5)
	bt.Add(11)
	bt.Add(6)
	bt.Print()

	//查找
	index, err := bt.Find(5)
	fmt.Println(index, err)
	index, err = bt.Find(2)
	fmt.Println(index, err)

	// 删除
	fmt.Println(bt.Delete(3))
	bt.Print()
	fmt.Println(bt.Delete(2))
	bt.Print()
}
