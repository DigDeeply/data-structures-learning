package binarytree

import "fmt"

// Bt 结构体
type Bt struct {
	Node []int
}

// New 初始化
func New(size int) *Bt {
	n := make([]int, size)
	return &Bt{Node: n}
}

// Add 添加节点
func (bt *Bt) Add(i int) {
	if bt.Node[1] == 0 {
		bt.Node[1] = i
		return
	}
	for index := 1; ; {
		if i < bt.Node[index] {
			if bt.Node[index*2] == 0 {
				bt.Node[index*2] = i
				break
			} else {
				index = index * 2
			}
		} else {
			if bt.Node[index*2+1] == 0 {
				bt.Node[index*2+1] = i
				break
			} else {
				index = index*2 + 1
			}
		}
	}
}

// Print 打印内容
func (bt *Bt) Print() {
	fmt.Println(bt.Node)
}

// Find 查找
func (bt *Bt) Find(i int) (index int, err error) {
	for index = 1; ; {
		if i < bt.Node[index] {
			if bt.Node[index*2] == 0 {
				break
			} else {
				index = index * 2
			}
		} else if i > bt.Node[index] {
			if bt.Node[index*2+1] == 0 {
				break
			} else {
				index = index*2 + 1
			}
		} else {
			return
		}
	}
	return 0, fmt.Errorf("not found the value %d", i)
}

// Delete 删除节点
func (bt *Bt) Delete(i int) (err error) {
	index, err := bt.Find(i)
	if err != nil {
		return err
	}
	bt.Node[index] = 0
	for bt.Node[index*2] != 0 {
		bt.Node[index], bt.Node[index*2] = bt.Node[index*2], bt.Node[index]
		index = index * 2
	}
	return
}
