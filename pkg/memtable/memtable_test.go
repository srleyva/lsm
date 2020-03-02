package memtable_test

import (
	"testing"

	. "github.com/srleyva/LSM/pkg/memtable"
	"github.com/srleyva/LSM/pkg/trees"
)

func TestMemtable(t *testing.T) {
	_, err := NewMemtable(10, trees.TREE_TYPE_RB)
	if err != nil {
		t.Errorf("err initializing memtable: %s", err)
	}

}
