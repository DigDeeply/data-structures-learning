package array

import (
	"fmt"
)

// 实现一个大小固定的有序数组，支持动态增删改操作
// 练习题：https://time.geekbang.org/column/article/80456

// Arr 数据结构体
type Arr struct {
	val       []int
	len       int
	maxLength int
}

// NewArr 新建Arr结构体
func NewArr(len int) *Arr {
	val := make([]int, len)
	arr := &Arr{
		val:       val,
		len:       0,
		maxLength: len,
	}
	return arr
}

// find 查找元素(递归) 即使没有找到元素，也返回最后的索引位置，用于方便插入数据
func (a *Arr) find(v, start, end int) (index int, s, e int, err error) {
	// fmt.Printf("start find %d from %d to %d in %v\n", v, start, end, a.val)
	if end == start {
		if a.val[start] == v {
			return start, 0, 0, nil
		}
		return 0, start, end, fmt.Errorf("not find the element: %d", v)
	}
	if end-start == 1 {
		if a.val[start] == v {
			return start, 0, 0, nil
		}
		if a.val[end] == v {
			return end, 0, 0, nil
		}
		return 0, start, end, fmt.Errorf("not find the element: %d", v)
	}
	m := start + (end+1-start)/2
	if a.val[m] > v {
		return a.find(v, start, m)
	}
	if a.val[m] < v {
		return a.find(v, m, end)
	}
	return m, 0, 0, nil
}

// Insert 插入元素
func (a *Arr) Insert(v int) (err error) {
	if a.len+1 > a.maxLength {
		return fmt.Errorf("array is full")
	}
	if a.len == 0 {
		a.val[0] = v
		a.len++
		return
	}

	index, s, _, err := a.find(v, 0, a.len-1)
	fmt.Printf("index:%d, start:%d, value:%d, %v\n", index, s, v, a.val)
	i := 0
	if err == nil {
		//存在相同的元素
		i = index
	} else {
		if a.val[s] > v {
			i = s
		} else {
			i = s + 1
		}
	}
	for j := a.len; j > i; j-- {
		a.val[j] = a.val[j-1]
	}
	a.val[i] = v
	a.len++
	return nil
}

// Delete 删除元素
func (a *Arr) Delete(v int) (err error) {
	index, _, _, err := a.find(v, 0, a.len-1)
	fmt.Printf("delete find %d in index=%d, %v, length:%d\n", v, index, a.val, a.len)
	if err != nil {
		return fmt.Errorf("not found the element: %d", v)
	}
	for i := index; i <= a.len-1; i++ {
		if i+1 >= a.len {
			a.val[i] = 0
			continue
		}
		a.val[i] = a.val[i+1]
	}
	a.len--
	return nil
}

// Print 打印元素
func (a *Arr) Print() {
	fmt.Printf("len: %d, maxLength: %d, value: %v\n", a.len, a.maxLength, a.val)
}
