package main

import (
	"fmt"
)

type cache struct {
	length     int
	currLength int
	service    *linked
}

type linked struct {
	key  string
	val  string
	next *linked
}

func main() {
	// 手写LRU算法
	fmt.Println("hello world.")

	//使用LRU算法
	c := &cache{}
	c.new(5)
	c.set("hello1", "world1")
	c.set("hello2", "world2")
	c.set("hello3", "world3")
	c.set("hello2", "world2")
	cacheVal, err := c.get("hello1")
	fmt.Println(cacheVal, err)
	c.debug()

}

func (c *cache) new(length int) {
	c.length = length
	c.currLength = 0
}

func (c *cache) set(key, val string) {
	node := &linked{
		key: key,
		val: val,
	}
	//如果没有元素
	if c.currLength == 0 {
		c.service = node
		c.currLength++
		return
	}
	head := c.service
	curr := c.service
	prev := c.service
	flag := false
	for i := 1; i <= c.currLength; i++ {
		if curr.key == key {
			//如果key已经存在节点里，挪动位置
			flag = true
			if head == curr {
				break
			}
			prev.next = curr.next
			node.next = head
			c.service = node
			break
		}
		prev = curr
		curr = curr.next
	}

	//如果key不存在，插入; 若已满删除一个旧节点
	if !flag {
		fmt.Println(c.currLength)
		if c.currLength >= c.length {
			//如果已满
			prev.next = nil
		} else {
			c.currLength++
		}
		// 插入到头部
		node.next = c.service
		c.service = node
		return
	}

}

func (c *cache) get(key string) (val string, err error) {
	//如果没有元素
	if c.currLength == 0 {
		return "", fmt.Errorf("key not exists")
	}
	//正常情况
	i := 1
	node := c.service
	for {
		if node.key == key {
			c.set(key, node.val)
			return node.val, nil
		}
		i++
		if i > c.currLength {
			return "", fmt.Errorf("key not exists")
		}
		node = node.next
	}
}

func (c *cache) debug() {
	// 顺序打印出所有节点
	fmt.Println("debug print Start.")
	node := c.service
	for i := 1; i <= c.currLength; i++ {
		fmt.Println(node.key, node.val)
		node = node.next
	}
	fmt.Println("debug print Done.")
}
