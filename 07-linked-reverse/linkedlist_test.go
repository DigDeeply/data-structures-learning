package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试链表
func TestLinkedList(t *testing.T) {
	tests := []struct {
		nodes       []*LinkedNode
		linkedPrint string
		err         error
	}{
		{
			nodes:       []*LinkedNode{},
			linkedPrint: "",
			err:         fmt.Errorf("empty linkedlist")},
		{
			nodes: []*LinkedNode{
				&LinkedNode{value: "val1"},
			},
			linkedPrint: "val1->nil",
			err:         nil},
		{
			nodes: []*LinkedNode{
				&LinkedNode{value: "val1"},
				&LinkedNode{value: "val2"},
			},
			linkedPrint: "val2->val1->nil",
			err:         nil},
		{
			nodes: []*LinkedNode{
				&LinkedNode{value: "val1"},
				&LinkedNode{value: "val2"},
				&LinkedNode{value: "val3"},
			},
			linkedPrint: "val3->val2->val1->nil",
			err:         nil},
	}
	for _, test := range tests {
		var prev *LinkedNode
		for _, node := range test.nodes {
			node.next = prev
			prev = node
		}
		ll := &LinkedList{head: &LinkedNode{next: prev}}
		msg, err := ll.Debug()
		assert.Equal(t, msg, test.linkedPrint)
		assert.Equal(t, err, test.err)
		ll.Print()
	}
}

// 测试反转
func TestReverse(t *testing.T) {
	tests := []struct {
		nodes       []*LinkedNode
		linkedPrint string
		err         error
	}{
		{
			nodes: []*LinkedNode{
				&LinkedNode{value: "val1"},
			},
			linkedPrint: "val1->nil",
			err:         nil},
		{
			nodes: []*LinkedNode{
				&LinkedNode{value: "val1"},
				&LinkedNode{value: "val2"},
			},
			linkedPrint: "val1->val2->nil",
			err:         nil},
		{
			nodes: []*LinkedNode{
				&LinkedNode{value: "val1"},
				&LinkedNode{value: "val2"},
				&LinkedNode{value: "val3"},
			},
			linkedPrint: "val1->val2->val3->nil",
			err:         nil},
		{
			nodes: []*LinkedNode{
				&LinkedNode{value: "val1"},
				&LinkedNode{value: "val2"},
				&LinkedNode{value: "val3"},
				&LinkedNode{value: "val4"},
			},
			linkedPrint: "val1->val2->val3->val4->nil",
			err:         nil},
	}
	for _, test := range tests {
		var prev *LinkedNode
		for _, node := range test.nodes {
			node.next = prev
			prev = node
		}
		ll := &LinkedList{head: &LinkedNode{next: prev}}
		//反转
		ll.Reverse()
		msg, err := ll.Debug()
		assert.Equal(t, msg, test.linkedPrint)
		assert.Equal(t, err, test.err)
	}
}

// TestHasCircle 检测是否有环
func TestHasCircle(t *testing.T) {
	ln1 := &LinkedNode{value: "val1"}
	ln2 := &LinkedNode{value: "val2"}
	ln3 := &LinkedNode{value: "val3"}
	ln4 := &LinkedNode{value: "val4"}
	ln5 := &LinkedNode{value: "val5"}
	tests := []struct {
		nodes     []*LinkedNode
		hasCircle bool
	}{
		{
			nodes: []*LinkedNode{
				ln1,
			},
			hasCircle: false},
		{
			nodes: []*LinkedNode{
				ln1, ln2,
			},
			hasCircle: false},
		{
			nodes: []*LinkedNode{
				ln1, ln2, ln1,
			},
			hasCircle: true},
		{
			nodes: []*LinkedNode{
				ln1, ln2, ln3, ln1,
			},
			hasCircle: true},
		{
			nodes: []*LinkedNode{
				ln1, ln2, ln3, ln4, ln1,
			},
			hasCircle: true},
		{
			nodes: []*LinkedNode{
				ln1, ln2, ln3, ln4, ln5, ln1,
			},
			hasCircle: true},
		{
			nodes: []*LinkedNode{
				ln1, ln2, ln3, ln4, ln2,
			},
			hasCircle: true},
		{
			nodes: []*LinkedNode{
				ln1, ln2, ln3, ln4, ln5, ln2,
			},
			hasCircle: true},
		{
			nodes: []*LinkedNode{
				ln1, ln2, ln3, ln4, ln5, ln3,
			},
			hasCircle: true},
	}
	for _, test := range tests {
		var prev *LinkedNode
		for _, node := range test.nodes {
			node.next = prev
			prev = node
		}
		ll := &LinkedList{head: &LinkedNode{next: prev}}
		//反转
		hasCircle := ll.HasCircle()
		assert.Equal(t, hasCircle, test.hasCircle)
	}
}

// TestHasCycle 优化版的有环检测
func TestHasCycle(t *testing.T) {
	ln1 := &LinkedNode{value: "val1"}
	ln2 := &LinkedNode{value: "val2"}
	ln3 := &LinkedNode{value: "val3"}
	ln4 := &LinkedNode{value: "val4"}
	ln5 := &LinkedNode{value: "val5"}
	tests := []struct {
		nodes     []*LinkedNode
		hasCircle bool
	}{
		{
			nodes: []*LinkedNode{
				ln1,
			},
			hasCircle: false},
		{
			nodes: []*LinkedNode{
				ln1, ln2,
			},
			hasCircle: false},
		{
			nodes: []*LinkedNode{
				ln1, ln2, ln1,
			},
			hasCircle: true},
		{
			nodes: []*LinkedNode{
				ln1, ln2, ln3, ln1,
			},
			hasCircle: true},
		{
			nodes: []*LinkedNode{
				ln1, ln2, ln3, ln4, ln1,
			},
			hasCircle: true},
		{
			nodes: []*LinkedNode{
				ln1, ln2, ln3, ln4, ln5, ln1,
			},
			hasCircle: true},
		{
			nodes: []*LinkedNode{
				ln1, ln2, ln3, ln4, ln2,
			},
			hasCircle: true},
		{
			nodes: []*LinkedNode{
				ln1, ln2, ln3, ln4, ln5, ln2,
			},
			hasCircle: true},
		{
			nodes: []*LinkedNode{
				ln1, ln2, ln3, ln4, ln5, ln3,
			},
			hasCircle: true},
	}
	for _, test := range tests {
		var prev *LinkedNode
		for _, node := range test.nodes {
			node.next = prev
			prev = node
		}
		ll := &LinkedList{head: &LinkedNode{next: prev}}
		//反转
		hasCircle := ll.HasCycle()
		assert.Equal(t, hasCircle, test.hasCircle)
	}
}
