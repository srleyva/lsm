package memtable

import (
	"fmt"

	"github.com/srleyva/LSM/pkg/trees"
)

type Memtable struct {
	tree     trees.Tree
	MemLimit int // MB
}

func NewMemtable(mbLimit int, treeType string) (*Memtable, error) {
	if mbLimit <= 0 {
		return nil, fmt.Errorf("invalid memory limit: %d", mbLimit)
	}

	tree, err := trees.Factory.NewTree(treeType)
	if err != nil {
		return nil, err
	}
	return &Memtable{
		tree:     tree,
		MemLimit: mbLimit,
	}, nil
}

func (m *Memtable) TreeSize() int {
	return m.tree.GetSize()
}

func (m *Memtable) Add(key string, value interface{}) error {
	m.tree.Add(key, value)
	if m.tree.GetSize() >= m.MemLimit {
		if err := m.FlushToSSTable(); err != nil {
			return err
		}
	}
	return nil
}

func (m *Memtable) Search(key string) (interface{}, error) {
	node, err := m.tree.Search(key)
	if err != nil {
		return nil, err
	}
	return node.GetValue(), nil
}

func (m *Memtable) FlushToSSTable() error {
	newTree, err := trees.Factory.NewTree(trees.TREE_TYPE_RB)
	if err != nil {
		return err
	}
	m.tree = newTree
	return nil
}
