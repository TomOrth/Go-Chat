//Package lists contains custom doubly linked lists that allow for fast insertion and deletion of elements
package lists

import "net"

//Type ConnList represents a list of client connections to a server
type ConnList struct {
	Head, Tail *Node //head and tail nodes, necessary for the list
	Size       int   //size of the list
}

//Append takes a client connection and appends it to the list
func (l *ConnList) Append(value net.Conn) {
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

//Delete takes a client connection and removes it from the list
func (l *ConnList) Delete(value net.Conn) {
	temp := l.Head
	for temp != nil {
		if temp.Value == value {
			if l.Head == temp {
				l.Head = temp.Next
			}
			if temp.Next != nil {
				temp.Next.Prev = temp.Prev
			}
			if temp.Prev != nil {
				temp.Prev.Next = temp.Next
			}
			l.Size -= 1
		}
		temp = temp.Next
	}
}

//Broadcast takes an incoming message and sends it to all connected clients
func (l *ConnList) Broadcast(message string) {
	temp := l.Head
	for temp != nil {
		temp.Value.(net.Conn).Write([]byte(message))
		temp = temp.Next
	}
}
