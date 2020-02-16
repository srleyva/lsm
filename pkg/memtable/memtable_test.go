package memtable_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/srleyva/LSM/pkg/memtable"
)

func TestMemtable(t *testing.T) {
	t.Run("test memtable add and get method", func(t *testing.T) {
		memtable, err := NewMemtable(15, "/tmp/memtable")
		defer os.RemoveAll("/tmp/memtable")
		if err != nil {
			t.Fatalf("err creating memtable: %s", err)
		}

		if err := memtable.Add("test", 12); err != nil {
			t.Errorf("err adding item to memtable: %s", err)
		}

		item := memtable.Get("test")
		if item == nil || item.(int) != 12 {
			t.Errorf("Expected: %d Actual: %d", 12, item.(int))
		}
	})
	t.Run("test flush to sstable", func(t *testing.T) {
		memtable, err := NewMemtable(1, "/tmp/memtable")
		defer os.RemoveAll("/tmp/memtable")
		if err != nil {
			t.Fatalf("err creating memtable: %s", err)
		}

		itemsToAdd := map[string]interface{}{
			"test":  121,
			"test1": 122,
			"test2": 123,
			"test3": 124,
			"test4": 125,
			"test5": 126,
			"test6": 127,
			"test7": 128,
			"test8": 129,
			"test9": 130,
		}

		for key, item := range itemsToAdd {
			if err := memtable.Add(key, item); err != nil {
				t.Errorf("err adding item to memtable: %s", err)
			}
		}

		matches, err := filepath.Glob("/tmp/memtable/sstable-*")
		if err != nil {
			t.Errorf("err getting sstables: %s", err)
		}

		if len(matches) <= 0 {
			t.Errorf("No sstables!")
		}
	})
}
