package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	count := 0
	currentNode := l.Head

	for currentNode != nil {
		count++
		currentNode = currentNode.Next
	}

	return count
}

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{
		Data: data,
		Next: nil,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		head := l.Head
		l.Head = newNode
		newNode.Next = head
	}
}
