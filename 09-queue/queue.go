package queue

import (
	"fmt"
)

type Node struct {
	val  interface{}
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func NewQueue() (q *Queue) {
	return &Queue{}
}

func (q *Queue) Print() {
	cur := q.head
	for cur != nil {
		fmt.Println(cur.val)
		if cur == q.tail {
			break
		}
		cur = cur.next
	}
}

func (q *Queue) Enqueue(val interface{}) {
	node := &Node{val: val}
	if q.head == nil {
		q.head = node
	}
	if q.tail == nil {
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
}

func (q *Queue) Dequeue() (val interface{}) {
	if q.head == nil {
		return nil
	}
	node := q.head
	q.head = q.head.next
	return node.val
}
