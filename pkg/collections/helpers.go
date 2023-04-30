package collections

type ListInterface interface {
	Remove(value interface{})
	Add(value interface{})
	Find(value interface{}) interface{}
	Display()
	Get(position int) interface{}
}

type SequencesInterface interface {
	Push(value interface{})
	Pop() interface{}
	Find(value interface{}) interface{}
	Display()
}

type DNode struct {
	prev  *DNode
	next  *DNode
	value interface{}
}

type SNode struct {
	value interface{}
	next  *SNode
}
