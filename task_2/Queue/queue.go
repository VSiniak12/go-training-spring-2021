package queue

import (
	"errors"
	"fmt"
	"reflect"
)

type Queue struct {
	first *Node
	last  *Node
	size  int
}

type Node struct {
	prev  *Node
	value interface{}
}

// define type string for struct Node
func (n Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

//Sort linkedlist in ascending order. You have got error if type of node value have not yet put in library
func (list *Queue) Sort() error {
	currentNode := list.last
	for currentNode != nil {
		if currentNode.prev != nil {
			i, err := compareValue(*currentNode)
			if err != nil {
				return err
			}
			if i {
				swap(currentNode)
				currentNode = list.last
			} else {
				currentNode = currentNode.prev
			}
		} else {
			break
		}
	}
	return nil
}

func swap(node *Node) {
	a1 := node.value
	b1 := node.prev.value
	node.value = b1
	node.prev.value = a1
}

func compareValue(node Node) (bool, error) {
	switch node.value.(type) {
	case int:
		return node.value.(int) > node.prev.value.(int), nil
	case string:
		return node.value.(string) > node.prev.value.(string), nil
	case rune:
		return node.value.(rune) > node.prev.value.(rune), nil
	case float64:
		return node.value.(float64) > node.prev.value.(float64), nil
	default:
		return false, errors.New("type of node value have not yet put in library")
	}
}

// define type string for struct Queue
func (q Queue) String() string {
	s := make([]interface{}, 0, 1)
	node := q.last
	for node != nil {
		s = append(s, node.value)
		node = node.prev
	}
	return fmt.Sprint(s)
}

//Add an element to the end of the queue. Shifts the element currently at the end (
//if any) to the left and increment size of queue
func (q *Queue) Enqueue(value interface{}) error {
	node := &Node{
		prev:  q.last,
		value: value,
	}
	if q.size != 0 && reflect.TypeOf(q.first.value) != reflect.TypeOf(node.value) {
		return fmt.Errorf("Invalid input value: %v", node.value)
	}
	q.last = node
	if q.size == 0 {
		q.first = node
	}
	q.size++
	return nil
}

//Displays the complete queue.
func (list *Queue) Display() {
	firstNode := list.last
	for firstNode != nil {
		fmt.Println(firstNode)
		firstNode = firstNode.prev
	}
}

//Remove an element from the front of the queue and decrement size of queue. If size of queue is equal 0,
//that you will get an error
func (q *Queue) Dequeue() error {
	if q.size == 0 {
		return errors.New("queue has not been initialized")
	}
	firstNode := q.last
	for i := q.size; i > 1; i-- {
		if i == 2 {
			firstNode.prev = nil
		}
		firstNode = firstNode.prev
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
func (q *Queue) Peek() (interface{}, error) {
	if q.size == 0 {
		return nil, errors.New("Queue have not yet initilization")
	}
	return q.first.value, nil
}

func Make(values ...interface{}) (*Queue, error) {
	q := Queue{}
	for _, v := range values {
		if err := q.Enqueue(v); err != nil {
			return nil, err
		}
	}
	return &q, nil
}
