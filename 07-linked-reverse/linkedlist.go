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

// Debug 方便调试
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

// Print 打印调试内容
func (ll *LinkedList) Print() {
	msg, err := ll.Debug()
	fmt.Println(msg, err)

}

// Reverse 单链表反转
// 使用一个临时的指针，缓存住下个节点的位置；交换当前位置的前后顺序，再用临时指针往下走。
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

// HasCircle 检测单链表是否有环
// 采用两步指针的情况，一个快指针，一个慢指针，快指针走每次走两步，慢指针每次走一步；如果快指针能跟慢指针重合，证明有环。
// 环存在有两种情况；一种是终点回到起点， 1->2->3->1； 另外一种是从终点回到中间节点，1->2->3->4->2;
// 下边的写法比较丑陋，时间复杂度为 O(3n),优雅的写法见 HasCycle
func (ll *LinkedList) HasCircle() (hasCircle bool) {

	if ll.head == nil || ll.head.next == nil || ll.head.next.next == nil {
		return false
	}
	fast := ll.head.next.next
	slow := ll.head.next
	for fast != nil {
		fast = fast.next
		if fast == nil {
			return false
		}
		if fast == slow {
			return true
		}
		fast = fast.next
		if fast == nil {
			return false
		}
		if fast == slow {
			return true
		}
		slow = slow.next
		if fast == slow {
			return true
		}
	}
	return false
}

// HasCycle 检测是否有环，时间复杂度 O(n)
func (ll *LinkedList) HasCycle() (hasCircle bool) {

	if ll.head == nil || ll.head.next == nil || ll.head.next.next == nil {
		return false
	}
	fast := ll.head.next
	slow := ll.head.next
	for fast != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			return true
		}
	}
	return false
}
