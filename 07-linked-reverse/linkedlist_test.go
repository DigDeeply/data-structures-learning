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
