package LinkedList

import (
	"fmt"
	"strings"
)

type node struct {
	element interface{}
	next    *node
}

type LinkedList struct {
	head *node
	size int
}

// Returns true if and only if the given index is valid.
func (lst *LinkedList) checkIndex(index int) bool {
	return index >= 0 && index < lst.size
}

// Adds an element at the end of the list.
func (lst *LinkedList) Add(element interface{}) {
	if lst.size == 0 {
		// If there are no elements in the list, simply add the new
		// element on the head.
		lst.head = &node{element, nil}
	} else {
		lastNode, _ := lst.getNode(lst.size - 1)
		lastNode.next = &node{element, nil}
	}
	lst.size++
}

// Returns the element at the given index.
func (lst *LinkedList) Get(index int) (interface{}, error) {
	node, err := lst.getNode(index)
	if err != nil {
		return nil, err
	}
	return node.element, nil
}

// Removes the element at the given index.
func (lst *LinkedList) Remove(index int) (oldVal interface{}, err error) {
	node, err := lst.getNode(index)
	if err != nil {
		return nil, err
	}
	oldVal = node.element
	err = nil
	if index == 0 {
		lst.head = node.next
		lst.size--
		return
	}
	previous, _ := lst.getNode(index - 1)
	previous.next = node.next
	lst.size--
	return
}

// Returns the node object at the given index. Only used internally.
func (lst *LinkedList) getNode(index int) (*node, error) {
	if !lst.checkIndex(index) {
		return nil, &Error{"Invalid index"}
	}
	curr := lst.head
	i := 0
	for ; i < index; i++ {
		curr = curr.next
	}
	return curr, nil
}

// Returns a string representation of the list.
func (lst LinkedList) String() string {
	curr := lst.head
	elements := make([]string, lst.size)
	for i := 0; curr != nil; i++ {
		elements[i] = fmt.Sprintf("%v", curr.element)
		curr = curr.next
	}
	return strings.Join(elements, ", ")
}
