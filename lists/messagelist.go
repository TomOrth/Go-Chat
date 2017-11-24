//Package lists contains custom doubly linked lists that allow for fast insertion and deletion of elements
package lists

//Type MsgList is a list of messages a client recieves from the server
type MsgList struct {
	Head, Tail *Node //head and tail nodes, necessary for the list
	Size       int   //size of the list
}

//Append takes a new msg as a string and appends it to the end of the list
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

//DeleteHead deletes the first message in the list
func (l *MsgList) DeleteHead() {
	l.Head = l.Head.Next
}

//MessageArr converts the doubly linked list into a string slice for the terminal UI
func (l *MsgList) MessageArr() []string {
	temp := make([]string, 0)
	tempNode := l.Head
	for tempNode != nil {
		temp = append(temp, tempNode.Value.(string))
		tempNode = tempNode.Next
	}
	return temp
}
