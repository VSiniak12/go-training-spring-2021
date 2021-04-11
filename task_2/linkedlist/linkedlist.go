package linkedlist

import (
	"errors"
	"fmt"
	"reflect"
)

type LinkedList struct {
	first *Node
	last  *Node
	size  int
}

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

// define type string for struct Node
func (n Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

// define type string for struct LinkedList
func (l LinkedList) String() string {
	s := make([]interface{}, 0, 1)
	f := l.first
	for f != nil {
		s = append(s, f.value)
		f = f.next
	}
	return fmt.Sprint(s)
}

//Inserts an element at the beginning of the list. Shifts the element currently at the beginning (
//if any) to the right and increment size of list
func (list *LinkedList) Insert(value interface{}) error {
	node := &Node{
		next:  list.first,
		value: value,
	}
	if list.size != 0 && reflect.TypeOf(list.first.value) != reflect.TypeOf(node.value) {
		return fmt.Errorf("Invalid input value: %v", node.value)
	}

	if list.first != nil {
		list.first.prev = node
	}
	list.first = node
	for node.next != nil {
		node = node.next
	}
	list.last = node
	list.size++
	return nil
}

//Deletes an element at the beginning of the list and decrement size of list. If size of list is equal 0,
//that you will get an error
func (list *LinkedList) Deletion() error {
	if list.size == 0 {
		return errors.New("list has not been initialized")
	}
	node := list.first.next
	list.first = node
	if list.size != 1 {
		list.first.prev = nil
	}
	list.size--
	return nil
}

//Displays the complete list.
func (list *LinkedList) Display() {
	firstNode := list.first
	for firstNode != nil {
		fmt.Println(firstNode)
		firstNode = firstNode.next
	}
}

//Searches an element using the id. If size of list is less than received index and index is less 0,
//that you will get an error
func (list *LinkedList) Search(index int) (*Node, error) {
	if index < 0 || index > list.size-1 {
		return nil, fmt.Errorf("Index equal %d and it out of list", index)
	}
	node := list.first
	for i := 0; i != index; i++ {
		node = node.next
	}
	return node, nil
}

//Deletes an element using the id. If size of list is less than received index and index is less 0,
//that you will get an error. Decrement size of list
func (list *LinkedList) Delete(index int) error {
	node, err := list.Search(index)
	if err != nil {
		return err
	}
	if node.prev == nil {
		node.next.prev = nil
		list.first = node.next
	} else if node.next == nil {
		node.prev.next = nil
		list.last = node.prev
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	list.size--
	return nil
}

//Sort linkedlist in ascending order. You have got error if type of node value have not yet put in library
func (list *LinkedList) Sort() error {
	currentNode := list.first
	for currentNode != nil {
		if currentNode.next != nil {
			currentNodeMoreNext, err := compareValue(*currentNode)
			if err != nil {
				return err
			}
			if currentNodeMoreNext {
				swap(currentNode)
				currentNode = list.first
			} else {
				currentNode = currentNode.next
			}
		} else {
			break
		}
	}
	return nil
}

func swap(node *Node) {
	a1 := node.value
	b1 := node.next.value
	node.value = b1
	node.next.value = a1
}

func compareValue(node Node) (bool, error) {
	switch node.value.(type) {
	case int:
		return node.value.(int) > node.next.value.(int), nil
	case string:
		return node.value.(string) > node.next.value.(string), nil
	case rune:
		return node.value.(rune) > node.next.value.(rune), nil
	case float64:
		return node.value.(float64) > node.next.value.(float64), nil
	default:
		return false, errors.New("type of node value have not yet put in library")
	}
}

//Make and initilization a linkedlist. You have got the error, if type of input value not equal first node.
func Make(values ...interface{}) (*LinkedList, error) {
	l := LinkedList{}
	for _, v := range values {
		if err := l.Insert(v); err != nil {
			return nil, err
		}
	}
	return &l, nil
}
