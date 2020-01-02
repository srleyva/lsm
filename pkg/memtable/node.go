package memtable

const (
	RED = iota
	BLACK
)

type Node struct {
	Color  int
	Key    string
	Value  interface{}
	Parent *Node
	Left   *Node
	Right  *Node
}

var nilNode = &Node{Color: BLACK}

func NewNode(key string, value interface{}) *Node {
	return &Node{
		Color:  BLACK,
		Parent: nil,
		Key:    key,
		Value:  value,
		Left:   nilNode,
		Right:  nilNode,
	}
}
