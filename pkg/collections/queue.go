package collections

import "fmt"

type Queue struct {
	head *DNode
	tail *DNode
}

func (Q *Queue) Push(value interface{}) {
	node := &DNode{
		next:  Q.head,
		value: value,
	}
	if Q.head != nil {
		Q.head.prev = node
	}
	Q.head = node

	if Q.tail == nil {
		Q.tail = node
	}
}

func (Q *Queue) Pop() interface{} {
	if Q.head == nil {
		return nil
	}

	popNode := Q.tail

	if Q.tail.prev == nil {
		Q.tail = nil
		Q.head = nil
		return popNode.value
	}
	Q.tail = Q.tail.prev
	Q.tail.next = nil
	return popNode.value
}

func (Q *Queue) Find(value interface{}) interface{} {
	if Q.head == nil {
		return nil
	}

	if Q.head.value == value {
		return Q.head.value
	}

	if Q.tail.value == value {
		return Q.tail.value
	}

	for current := Q.head; current != nil; current = current.next {
		if current.value == value {
			return current.value
		}
	}

	return nil
}

func (Q *Queue) Display() {
	current := Q.head
	for current != nil {
		fmt.Printf("%v\n", current.value)
		current = current.next
	}
}
