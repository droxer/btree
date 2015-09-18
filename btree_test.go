package btree_test

import (
	"github.com/droxer/btree"
	"io"
	"os"
	"testing"
)

type item int

func (i item) Less(other btree.Item) bool {
	return i < other.(item)
}

func TestNotGet(t *testing.T) {
	bt := btree.New(2)
	bt.Insert(item(1))
	bt.Insert(item(2))
	bt.Insert(item(3))
	bt.Insert(item(4))
	bt.Insert(item(5))
	bt.Insert(item(8))
	bt.Insert(item(9))

	if bt.Get(item(7)) != nil {
		t.Fatalf("expected is nil, actual is %v", bt.Get(item(7)))
	}
}

func TestGet(t *testing.T) {
	bt := btree.New(2)
	bt.Insert(item(1))
	bt.Insert(item(2))
	bt.Insert(item(3))
	bt.Insert(item(4))
	bt.Insert(item(5))
	bt.Insert(item(8))
	bt.Insert(item(9))

	if bt.Get(item(8)) != item(8) {
		t.Fatalf("expected is 5, actual is %v", bt.Get(item(7)))
	}
}

func ExampleInsert() {
	bt := btree.New(2)
	bt.Insert(item(1))
	bt.Insert(item(2))
	bt.Insert(item(3))
	bt.Insert(item(4))
	bt.Insert(item(5))
	bt.Insert(item(8))
	bt.Insert(item(9))

	bt.Print(os.Stdout)
	// Output: 2--4
	//     1    3    5--8--9
}

func ExampleDelete() {
	bt := btree.New(2)
	bt.Insert(item(1))
	bt.Insert(item(2))
	bt.Insert(item(3))
	bt.Insert(item(4))
	bt.Insert(item(5))
	bt.Insert(item(8))
	bt.Insert(item(9))

	bt.Delete(item(5))
	bt.Print(os.Stdout)
	// Output: 2--4
	//     1    3    8--9
}
