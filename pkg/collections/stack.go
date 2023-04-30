package collections

import "fmt"

type Stack struct {
	head *SNode
}

func (S *Stack) Push(value interface{}) {
	node := &SNode{
		next:  S.head,
		value: value,
	}
	S.head = node
}

func (S *Stack) Pop() interface{} {
	if S.head == nil {
		return nil
	}
	popElement := S.head
	S.head = S.head.next
	return popElement.value
}

func (S *Stack) Find(value interface{}) interface{} {
	if S.head == nil {
		return nil
	}

	if S.head.value == value {
		return S.head.value
	}

	for current := S.head; current != nil; current = current.next {
		if current.value == value {
			return current.value
		}
	}

	return nil
}

func (S *Stack) Display() {
	current := S.head
	for current != nil {
		fmt.Printf("%v\n", current.value)
		current = current.next
	}
}
