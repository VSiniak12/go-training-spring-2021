package main

import (
	"errors"
	"fmt"
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

//Inserts an element at the beginning of the list. Shifts the element currently at the beginning (
//if any) to the right and increment size of list
func (list *LinkedList) Insert(value interface{}) {
	node := &Node{
		next:  list.first,
		value: value,
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
		fmt.Println(firstNode.value)
		firstNode = firstNode.next
	}
}

//Searches an element using the id. If size of list is less than received index and index is less 0,
//that you will get an error
func (list *LinkedList) Search(index int) (*Node, error) {
	if index < 0 || index > list.size-1 {
		return nil, fmt.Errorf("Index is %d and it out of list", index)
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

func (list *LinkedList) Sort() {

}

func main() {
	linkedList := LinkedList{}
	/*linkedList.Insert("1")
	linkedList.Insert("1")
	linkedList.Insert("2")
	linkedList.Insert("3")*/
	/*linkedList.Insert(1)
	linkedList.Insert(6)
	linkedList.Insert(22)
	linkedList.Insert(3)
	linkedList.Insert(13)
	linkedList.Insert(22)
	linkedList.Insert(11)
	linkedList.Insert(65)*/
	linkedList.Display()
	err := linkedList.Deletion()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("after deletion")
	linkedList.Display()
	err = linkedList.Delete(-5)
	if err != nil {
		fmt.Println(err)
	}
	linkedList.Display()
	linkedList.Sort()
}
