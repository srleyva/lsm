package trees

import (
	"fmt"
)

type Node interface {
	GetKey() string
	GetValue() interface{}
}

type Tree interface {
	GetType() string
	GetSize() int
	Add(string, interface{})
	Search(string) (Node, error)
	ToList() []Node
}

var Factory *factory = &factory{
	types: map[string]func() Tree{},
}

type factory struct {
	types map[string]func() Tree
}

func (f *factory) NewTree(treeType string) (Tree, error) {
	newTreeFunc, ok := f.types[treeType]
	if !ok {
		return nil, fmt.Errorf("err creating tree: %s not found in factory", treeType)
	}
	return newTreeFunc(), nil
}

func (f *factory) Register(name string, constructor func() Tree) {
	f.types[name] = constructor
}
