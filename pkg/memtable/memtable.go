package memtable

import "fmt"

type Memtable struct {
	tree     *RBTree
	MemLimit int // MB
}

func NewMemtable(mbLimit int) (*Memtable, error) {
	if mbLimit <= 0 {
		return nil, fmt.Errorf("invalid memory limit: %d", mbLimit)
	}
	return &Memtable{
		tree:     NewRBTree(),
		MemLimit: mbLimit,
	}, nil
}

func (m *Memtable) Add(key string, value interface{}) error {

	if m.tree.Size >= m.MemLimit {
		if err := m.FlushToSSTable(); err != nil {
			return err
		}
	}
	return nil
}

func (m *Memtable) FlushToSSTable() error {
	return nil
}
