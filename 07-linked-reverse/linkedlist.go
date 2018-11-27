package linkedlist

import (
	"fmt"
)

type LinkedNode struct {
	value interface{}
	next  *LinkedNode
}

type LinkedList struct {
	head *LinkedNode
}

// 方便调试
func (ll *LinkedList) Debug() (msg string, err error) {
	if ll.head == nil || ll.head.next == nil {
		return "", fmt.Errorf("empty linkedlist")
	}
	for cur := ll.head.next; cur != nil; cur = cur.next {
		msg += fmt.Sprintf("%v->", cur.value)
	}
	msg += fmt.Sprint("nil")
	return msg, nil
}

// 打印调试内容
func (ll *LinkedList) Print() {
	msg, err := ll.Debug()
	fmt.Println(msg, err)

}

// 单链表反转
func (ll *LinkedList) Reverse() {
	if ll.head == nil || ll.head.next == nil || ll.head.next.next == nil {
		fmt.Println("no need reverse")
	}
	cur := ll.head.next
	var prev *LinkedNode
	for cur != nil {
		tmp := cur.next
		cur.next = prev
		prev = cur
		cur = tmp
	}
	ll.head.next = prev
}
