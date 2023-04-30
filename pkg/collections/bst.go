package collections

import (
	"bytes"
	"fmt"
)

type Bst struct {
	value []byte
	left  *Bst
	right *Bst
}

func (node *Bst) Insert(value []byte) {
	if bytes.Compare(value, node.value) < 0 {
		if node.left == nil {
			node.left = &Bst{
				left:  nil,
				right: nil,
				value: value,
			}
		} else {
			node.left.Insert(value)
		}
	} else {
		if node.right == nil {
			node.right = &Bst{
				left:  nil,
				right: nil,
				value: value,
			}
		} else {
			node.right.Insert(value)
		}
	}
}

func (node *Bst) Find(value []byte) bool {
	if node == nil {
		return false
	}
	comparator := bytes.Compare(node.value, value)
	if comparator == -1 {
		return node.right.Find(value)
	} else if comparator == 1 {
		return node.left.Find(value)
	} else {
		return true
	}
}

func (node *Bst) Display() {
	if node.left != nil {
		node.left.Display()
	}
	fmt.Printf("%v\n", string(node.value))
	if node.right != nil {
		node.right.Display()
	}
}

func (node *Bst) DisplaySymmetrical() {
	fmt.Printf("%v\n", string(node.value))
	if node.left != nil {
		node.left.Display()
	}
	if node.right != nil {
		node.right.Display()
	}
}

func (node *Bst) DisplayReverso() {
	if node.right != nil {
		node.right.DisplayReverso()
	}
	fmt.Printf("%v\n", string(node.value))
	if node.left != nil {
		node.left.DisplayReverso()
	}
}
