package utils

type Node struct {
    Value User
	Next, Prev *Node
}
type List struct {
	Head, Tail *Node
}

func (l *List) Append(value User) {
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
}

func (l *List) Prepend(value User) {
	node := &Node{value, nil, nil}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
}

func (l *List) Get(index int) User {
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