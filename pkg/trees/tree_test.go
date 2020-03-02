package trees_test

import (
	"testing"

	. "github.com/srleyva/LSM/pkg/trees"
)

func TestFactory(t *testing.T) {
	t.Run("test provisioning of tree", func(t *testing.T) {
		tree, err := Factory.NewTree(TREE_TYPE_RB)
		if err != nil {
			t.Errorf("err returned where not expected: %s", err)
		}

		if tree.GetType() != TREE_TYPE_RB {
			t.Errorf("wrong type of tree returned from factory: %s", tree.GetType())
		}
	})
}
