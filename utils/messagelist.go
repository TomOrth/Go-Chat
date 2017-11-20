package utils

type MsgList struct {
	Head, Tail *Node
	Size       int
}

func (l *MsgList) Append(value string) {
	node := &Node{value, nil, nil}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		node.Prev = l.Tail
		l.Tail = node
		if l.Head.Next == nil {
			l.Head.Next = l.Tail.Prev
		}
	}
	l.Size += 1
}

func (l *MsgList) Prepend(value string) {
	node := &Node{value, nil, nil}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
	l.Size += 1
}

func (l *MsgList) Get(index int) string {
	temp := l.Head
	counter := 0
	for temp != nil {
		if counter == index {
			return temp.Value.(string)
		}
		temp = temp.Next
		counter += 1
	}
	return ""
}

func (l *MsgList) DeleteHead() {
	l.Head = l.Head.Next
}

func (l *MsgList) MessageArr() []string {
	temp := make([]string, 0)
	tempNode := l.Head
	for tempNode != nil {
		temp = append(temp, tempNode.Value.(string))
		tempNode = tempNode.Next
	}
	return temp
}
