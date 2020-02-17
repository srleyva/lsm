package memtable

import (
	"unsafe"
)

type RBTree struct {
	Root *Node
	Size int
	list []Item
}

func NewRBTree() *RBTree {
	return &RBTree{
		Root: nil,
	}
}

type Item struct {
	Key     string
	Value   interface{}
	Deleted bool
}

func (r *RBTree) Add(key string, value interface{}) {

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

func (r *RBTree) fixTreeAfterAdd(newNode *Node) {
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
					r.RotateLeft(newNode)
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
				r.RotateLeft(newNode.Parent.Parent)
			}
		}
	}
	r.Root.Color = BLACK
}
func (r *RBTree) RotateLeft(newNode *Node) {
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
func (r *RBTree) rotateRight(newNode *Node) {
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
func (r *RBTree) Search(key string) *Node {
	currentNode := r.Root
	for currentNode != NilNode && key != currentNode.Key {
		if key < currentNode.Key {
			currentNode = currentNode.Left
		}

		if key > currentNode.Key {
			currentNode = currentNode.Right
		}
	}
	return currentNode
}

// ToList returns a inorder list of tree
func (r *RBTree) ToList() []Item {
	// InOrder traversal to serialize nodes to list
	r.inOrderTraverse(r.Root)
	return r.list
}

func (r *RBTree) inOrderTraverse(node *Node) {
	if node.Left != NilNode {
		r.inOrderTraverse(node.Left)
	}
	r.list = append(r.list, Item{node.Key, node.Value, false})
	if node.Right != NilNode {
		r.inOrderTraverse(node.Right)
	}
}
