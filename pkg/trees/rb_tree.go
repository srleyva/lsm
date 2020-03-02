package trees

import (
	"fmt"
	"sync"
	"unsafe"
)

const TREE_TYPE_RB = "redblack"

type Color uint

const (
	RED Color = iota
	BLACK
)

type RedBlackNode struct {
	Color  Color
	Key    string
	Value  interface{}
	Parent *RedBlackNode
	Left   *RedBlackNode
	Right  *RedBlackNode
}

func (r *RedBlackNode) GetKey() string {
	return r.Key
}

func (r *RedBlackNode) GetValue() interface{} {
	return r.Value
}

var NilNode = &RedBlackNode{Color: BLACK}

func NewNode(key string, value interface{}) *RedBlackNode {
	return &RedBlackNode{
		Color:  BLACK,
		Parent: nil,
		Key:    key,
		Value:  value,
		Left:   NilNode,
		Right:  NilNode,
	}
}

type RBTree struct {
	Type string
	Root *RedBlackNode
	Size int

	sync.Mutex
}

func NewRBTree() Tree {
	return &RBTree{
		Type: TREE_TYPE_RB,
		Root: nil,
	}
}

func init() {
	Factory.Register(TREE_TYPE_RB, NewRBTree)
}

func (r *RBTree) GetType() string {
	return r.Type
}

func (r *RBTree) GetSize() int {
	return r.Size
}

func (r *RBTree) Add(key string, value interface{}) {
	r.Lock()
	defer r.Unlock()
	newNode := NewNode(key, value)
	payloadSize := int(unsafe.Sizeof(newNode))

	if r.Root == nil {
		r.Root = newNode
		r.Size += payloadSize
	} else {

		currentNode := r.Root
		for {
			if newNode.Key == currentNode.Key {
				currentNode.Value = newNode.Value
				return
			}
			if newNode.Key > currentNode.Key {
				if currentNode.Right == NilNode {
					newNode.Parent = currentNode
					currentNode.Right = newNode
					break
				}
				currentNode = currentNode.Right
			}
			if newNode.Key < currentNode.Key {
				if currentNode.Left == NilNode {
					newNode.Parent = currentNode
					currentNode.Left = newNode
					break
				}
				currentNode = currentNode.Left
			}
		}
		r.fixTreeAfterAdd(newNode)
	}
}

func (r *RBTree) fixTreeAfterAdd(newNode *RedBlackNode) {
	newNode.Color = RED
	for newNode != r.Root && newNode.Parent.Color == RED {
		if newNode.Parent.Parent.Left == newNode.Parent {
			uncle := newNode.Parent.Parent.Right
			if uncle.Color == RED {
				newNode.Parent.Color = BLACK
				uncle.Color = BLACK
				newNode.Parent.Parent.Color = RED
				newNode = newNode.Parent.Parent
			} else {
				if newNode == newNode.Parent.Right {
					newNode = newNode.Parent
					r.rotateLeft(newNode)
				}
				newNode.Parent.Color = BLACK
				newNode.Parent.Parent.Color = RED
				r.rotateRight(newNode.Parent.Parent)
			}
		} else if newNode.Parent.Parent.Right == newNode.Parent {
			// TODO newNode Parent is right child
			uncle := newNode.Parent.Parent.Left
			if uncle.Color == RED {
				newNode.Parent.Color = BLACK
				uncle.Color = BLACK
				newNode = newNode.Parent.Parent
			} else {
				if newNode == newNode.Parent.Left {
					newNode = newNode.Parent
					r.rotateRight(newNode)
				}
				newNode.Parent.Color = BLACK
				newNode.Parent.Parent.Color = RED
				r.rotateLeft(newNode.Parent.Parent)
			}
		}
	}
	r.Root.Color = BLACK
}
func (r *RBTree) rotateLeft(newNode *RedBlackNode) {
	RightNode := newNode.Right

	newNode.Right = RightNode.Left

	RightNode.Left.Parent = newNode

	RightNode.Parent = newNode.Parent

	// Correct top of tree

	// If this is the top of the tree
	// Make RightNode root
	// If This is not top of tree, need to fix newNode's parent node
	// Check if newNode is Left or Right of Parent Node and assign RightNode

	if RightNode.Parent == nil {
		r.Root = RightNode
	} else if newNode == newNode.Parent.Left {
		newNode.Parent.Left = RightNode
	} else if newNode == newNode.Parent.Right {
		newNode.Parent.Right = RightNode
	}

	RightNode.Left = newNode
	newNode.Parent = RightNode

}
func (r *RBTree) rotateRight(newNode *RedBlackNode) {
	LeftNode := newNode.Left
	newNode.Left = LeftNode.Right

	LeftNode.Right.Parent = newNode
	LeftNode.Parent = newNode.Parent

	if LeftNode.Parent == nil {
		r.Root = LeftNode
	} else if newNode == newNode.Parent.Left {
		newNode.Parent.Left = LeftNode
	} else if newNode == newNode.Parent.Right {
		newNode.Parent.Right = LeftNode
	}

	LeftNode.Right = newNode
	newNode.Parent = LeftNode
}

// Search will check if key is in tree and return if so
// If it does not exist it will return NilNode
func (r *RBTree) Search(key string) (Node, error) {
	r.Lock()
	defer r.Unlock()
	currentNode := r.Root
	for currentNode != NilNode && key != currentNode.Key {
		if key < currentNode.Key {
			currentNode = currentNode.Left
		}
		if key > currentNode.Key {
			currentNode = currentNode.Right
		}
	}

	if currentNode == NilNode {
		return nil, fmt.Errorf("key does not exist in tree: %s", key)
	}
	return currentNode, nil
}

// ToList returns a inorder list of tree
func (r *RBTree) ToList() []Node {
	r.Lock()
	defer r.Unlock()
	// InOrder traversal to serialize nodes to list
	return r.inOrderList()
}

func (r *RBTree) inOrderList() []Node {

	inorderList := []Node{}
	current := r.Root
	stack := []*RedBlackNode{current}

	for n := len(stack); n > 0; n = len(stack) {
		if current != NilNode {
			stack = append(stack, current)
			current = current.Left
			continue
		}

		current = stack[n-1]
		inorderList = append(inorderList, current)
		stack = stack[:n-1]
		current = current.Right
	}

	return inorderList[:len(inorderList)-1]
}
