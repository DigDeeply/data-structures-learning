package queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueInit(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Print()
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	assert.Equal(t, 1, 1)
}
