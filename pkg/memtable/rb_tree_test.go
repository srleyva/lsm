package memtable_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	. "github.com/srleyva/LSM/pkg/memtable"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func TestRBTress(t *testing.T) {
	t.Run("Test Add and Search in RBTree", func(t *testing.T) {
		rbTree := NewRBTree()

		rbTree.Add("test", 345)

		if rbTree.Root.Key != "test" || rbTree.Root.Value != 345 {
			t.Errorf("Wrong Values returned: expected=[%s]:[%d] got=[%s]:[%d]", "test", 345, rbTree.Root.Key, rbTree.Root.Value)
		}

		rbTree.Add("apple", 346)
		if rbTree.Root.Left.Key != "apple" || rbTree.Root.Left.Value != 346 {
			t.Errorf("Wrong Values returned: expected=[%s]:[%d] got=[%s]:[%d]", "test", 345, rbTree.Root.Left.Key, rbTree.Root.Left.Value)
		}

		rbTree.Add("zoo", 347)
		if rbTree.Root.Right.Key != "zoo" || rbTree.Root.Right.Value != 347 {
			t.Errorf("Wrong Values returned: expected=[%s]:[%d] got=[%s]:[%d]", "test", 345, rbTree.Root.Right.Key, rbTree.Root.Right.Value)
		}

		node := rbTree.Search("zoo")
		if node.Key != "zoo" || node.Value != 347 {
			t.Errorf("Search did not return expected result: expected=[%s][%d] actual=[%s][%d]", "zoo", 347, node.Key, node.Value)
		}
	})

	t.Run("Test balancing of RB Tree", func(t *testing.T) {
		entries := 100
		tree := NewRBTree()
		for i := 0; i < entries; i++ {
			tree.Add(String(10), nil)
		}

		maxHeight := 2 * LogN(entries+1)
		node := tree.Root
		count := 0
		for node != NilNode {
			count++
			node = node.Right
		}

		if count > maxHeight+1 {
			t.Errorf("Broke Level Contraints: got=[%d] expected=[less than %d]", count, maxHeight+1)
		}

	})

	t.Run("Test sort of RB Tree", func(t *testing.T) {
		entryCount := 20
		entries := make([]string, entryCount)
		for i := 0; i < entryCount; i++ {
			entries[i] = String(10)
		}

		tree := NewRBTree()
		for i, entry := range entries {
			tree.Add(entry, i)
		}

		sort.Strings(entries)

		list := tree.ToList()

		if len(list) != len(entries) {
			t.Errorf("Incorrect List Size got=[%d] expected=[%d]", len(list), len(entries))
		}

		for i, entry := range entries {
			if entry != list[i].Key {
				t.Errorf("Not sorted order: expected=[%s] got=[%s]", entry, list[i].Key)
			}
		}
	})

	t.Run("Test search of RB Tree", func(t *testing.T) {
		entryCount := 20
		entries := make([]string, entryCount)
		for i := 0; i < entryCount; i++ {
			entries[i] = String(10)
		}

		tree := NewRBTree()
		for i, entry := range entries {
			tree.Add(entry, i)
		}

		randomIndex := rand.Intn(entryCount)

		entry := tree.Search(entries[randomIndex])

		if entry.Value != randomIndex {
			t.Errorf("search returned incorrect value: got=[%d] expected=[%d]", entry.Value, randomIndex)
		}

	})
}

func LogN(n int) int {
	if n < 1 {
		return 0
	}

	return 1 + LogN(n/2)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
