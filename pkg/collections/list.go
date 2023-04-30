package collections

import "fmt"

type List struct {
	head *DNode
	tail *DNode
}

func (L *List) Add(value interface{}) {
	node := &DNode{
		next:  L.head,
		value: value,
	}
	if L.head != nil {
		L.head.prev = node
	}
	L.head = node

	if L.tail == nil {
		L.tail = node
	}
}

func (L *List) Remove(value interface{}) {
	if L.head == nil {
		return
	}

	if L.head.value == value {
		L.head = L.head.next
		return
	}

	if L.tail.value == value {
		L.tail = L.tail.prev
		L.tail.next = nil
		return
	}

	for current := L.head; current != nil; current = current.next {
		if current.value == value {
			next := current.next
			current = current.prev
			current.next = next
			return
		}
	}
}

func (L *List) Find(value interface{}) interface{} {
	if L.head == nil {
		return nil
	}

	if L.head.value == value {
		return L.head.value
	}

	if L.tail.value == value {
		return L.tail.value
	}

	for current := L.head; current != nil; current = current.next {
		if current.value == value {
			return current.value
		}
	}

	return nil
}

func (L *List) Get(position int) interface{} {
	if L.head == nil {
		return nil
	}

	current := L.head
	counter := 0
	for current != nil {
		if counter == position {
			return current.value
		}
		counter++
		current = current.next
	}
	return nil
}

func (L *List) Display() {
	current := L.head
	for current != nil {
		fmt.Printf("%v\n", current.value)
		current = current.next
	}
}
