package utils

type Node struct {
    Value User
	Next, Prev *Node
}
type ConnList struct {
	Head, Tail *Node
	Size int
}

func (l *ConnList) Append(value User) {
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

func (l *ConnList) Prepend(value User) {
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

func (l *ConnList) Get(index int) User {
	temp := l.Head
	counter := 0
	for temp != nil {
		if counter == index {
			return temp.Value
		}
		temp = temp.Next
		counter += 1;
	}
	return User{}
}

func (l *ConnList) Broadcast(message string) {
	temp := l.Head
	for temp != nil {
		temp.Value.Conn.Write([]byte(message))
		temp = temp.Next
	}
}