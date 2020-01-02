package memtable

import "unsafe"

type RBTree struct {
	Root *Node
	Size int
}

func NewRBTree() *RBTree {
	return &RBTree{
		Root: nilNode,
	}
}

func (r *RBTree) Add(key string, value interface{}) {
	newNode := NewNode(key, value)
	payloadSize := int(unsafe.Sizeof(key) + unsafe.Sizeof(key))
	if r.Root == nilNode {
		r.Root = newNode
		r.Root.Color = BLACK
		r.Size += payloadSize
		return
	}

	currentNode := r.Root
	var potentialParent *Node
	for currentNode != nilNode {
		potentialParent = currentNode
		if newNode.Key < currentNode.Key {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}

	newNode.Parent = potentialParent
	if newNode.Key < newNode.Parent.Key {
		newNode.Parent.Left = newNode
	} else {
		newNode.Parent.Right = newNode
	}
	r.FixTreeAfterAdd(newNode)
}

// Search will check if key is in tree and return if so
func (r *RBTree) Search(key string) *Node {
	currentNode := r.Root
	for currentNode != nilNode && key != currentNode.Key {
		if key < currentNode.Key {
			currentNode = currentNode.Left
		}

		if key > currentNode.Key {
			currentNode = currentNode.Right
		}
	}
	return currentNode
}

func (r *RBTree) FixTreeAfterAdd(newNode *Node) {}

func (r *RBTree) RotateLeft(newNode *Node) {}

func (r *RBTree) RotateRight(newNode *Node) {}

func (r *RBTree) GetAllNodes() map[string]interface{} {
	return nil
}

func (r *RBTree) IsRed() bool {
	return r.Root.Color == RED
}

func (r *RBTree) IsBlack() bool {
	return r.Root.Color == BLACK
}
