package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	first *Node
	last  *Node
	size  int
}

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

//Add an element to the end of the queue. Shifts the element currently at the end (
//if any) to the left and increment size of queue
func (q *Queue) Enqueue(value interface{}) {
	node := &Node{
		prev:  q.last,
		value: value,
	}
	if q.last != nil {
		q.last.next = node
	}
	q.last = node
	for node.prev != nil {
		node = node.prev
	}
	q.first = node
	q.size++
}

//Remove an element from the front of the queue and decrement size of queue. If size of queue is equal 0,
//that you will get an error
func (q *Queue) Dequeue() error {
	if q.size == 0 {
		return errors.New("queue has not been initialized")
	}
	node := q.first.next
	q.first = node
	if q.size != 1 {
		q.first.prev = nil
	}
	q.size--
	return nil
}

//Check if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

//Check if the queue is full
func (q *Queue) IsFull() bool {
	return q.size != 0
}

//Get the value of the front of the queue without removing it
func (q *Queue) Peek() interface{} {
	return q.first.value
}

func main() {
	q := Queue{}
	fmt.Println(q.IsEmpty())
	fmt.Println(q.IsFull())
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.IsFull())
	fmt.Println(q.Peek())
}
