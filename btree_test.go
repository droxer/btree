package btree

import (
	"testing"
)

type item int

func (i item) Less(other Item) bool {
	return i < other.(item)
}

func TestAddFirstNode(t *testing.T) {
	i := item(2)

	btree := New(2)
	btree.Insert(i)

	found := btree.root.items[0].(item)
	if found != i {
		t.Errorf("expected %s, actual is %s", i, found)
	}
}
