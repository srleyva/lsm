package trees_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	. "github.com/srleyva/LSM/pkg/trees"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func BenchmarkRBTree(b *testing.B) {
	entryCount := 10000
	entries := make([]string, entryCount)
	for i := 0; i < entryCount; i++ {
		entries[i] = String(10)
	}

	tree := NewRBTree().(*RBTree)
	for i, entry := range entries {
		tree.Add(entry, i)
	}

	b.Run("benchmark search", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			randomIndex := rand.Intn(entryCount)
			tree.Search(entries[randomIndex])
		}
	})

	b.Run("benchmark list", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			tree.ToList()
		}
	})

	b.Run("benchmark add", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			tree.Add(string(n), n)
		}
	})
}

func TestRBTress(t *testing.T) {
	// Generate List
	entryCount := 10000
	entries := make([]string, entryCount)
	for i := 0; i < entryCount; i++ {
		entries[i] = String(10)
	}

	t.Run("Test add override key", func(t *testing.T) {
		// Set Up tree
		tree := NewRBTree().(*RBTree)
		for i, entry := range entries {
			tree.Add(entry, i)
		}

		// Prepare test
		sleyva := struct {
			FirstName string
			LastName  string
			Employer  string
		}{
			"Stephen",
			"Leyva",
			"Blizzard",
		}
		tree.Add("sleyva", sleyva)
		sleyva.Employer = "Amazon Web Services"
		tree.Add("sleyva", sleyva)

		actual := tree.Search("sleyva").GetValue().(struct {
			FirstName string
			LastName  string
			Employer  string
		})

		if actual.Employer != sleyva.Employer {
			t.Errorf("Updated entry not correct: expected=%s got=%s", sleyva.Employer, actual.Employer)
		}
	})

	t.Run("Test balancing of RB Tree", func(t *testing.T) {

		// Set Up tree
		tree := NewRBTree().(*RBTree)
		for i, entry := range entries {
			tree.Add(entry, i)
		}
		maxHeight := 2 * LogN(entryCount+1)
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
		// Set Up tree
		tree := NewRBTree().(*RBTree)
		for i, entry := range entries {
			tree.Add(entry, i)
		}

		sort.Strings(entries)
		list := tree.ToList()
		if len(list) != len(entries) {
			t.Errorf("Incorrect List Size got=[%d] expected=[%d]", len(list), len(entries))
		}
		for i, item := range list {
			if item.GetKey() != entries[i] {
				t.Errorf("Not sorted order: expected=[%s] got=[%s]", entries[i], item.GetKey())
			}
		}
	})

	t.Run("Test search of RB Tree", func(t *testing.T) {
		// Set Up tree
		tree := NewRBTree().(*RBTree)
		for i, entry := range entries {
			tree.Add(entry, i)
		}

		randomIndex := rand.Intn(entryCount)

		entry := tree.Search(entries[randomIndex])

		if entry.GetValue() != randomIndex {
			t.Errorf("search returned incorrect value: got=[%d] expected=[%d]", entry.GetValue(), randomIndex)
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
