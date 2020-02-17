package memtable

type Color uint

const (
	RED Color = iota
	BLACK
)

type Node struct {
	Color  Color
	Key    string
	Value  interface{}
	Parent *Node
	Left   *Node
	Right  *Node
}

var NilNode = &Node{Color: BLACK}

func NewNode(key string, value interface{}) *Node {
	return &Node{
		Color:  BLACK,
		Parent: nil,
		Key:    key,
		Value:  value,
		Left:   NilNode,
		Right:  NilNode,
	}
}
