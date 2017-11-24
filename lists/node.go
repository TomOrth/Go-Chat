//Package lists contains custom doubly linked lists that allow for fast insertion and deletion of elements
package lists

//Type Node represents a node inside a list
type Node struct {
	Value      interface{} //generic value type
	Next, Prev *Node       //references to the next and prev nodes
}
