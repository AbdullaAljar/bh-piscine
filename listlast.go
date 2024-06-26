package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListLast(l *List) interface{} {
	if l.Tail != nil {
		return l.Tail.Data
	}
	return nil
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{
		Data: data,
		Next: nil,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
