package btree_test

import (
	"github.com/droxer/btree"
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
	bt.Insert(item(6))
	bt.Insert(item(7))
	bt.Insert(item(8))
	bt.Insert(item(9))
	bt.Insert(item(10))

	if bt.Get(item(11)) != nil {
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
	bt.Insert(item(6))
	bt.Insert(item(7))
	bt.Insert(item(8))
	bt.Insert(item(9))
	bt.Insert(item(10))

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
	bt.Insert(item(6))
	bt.Insert(item(7))
	bt.Insert(item(8))
	bt.Insert(item(9))
	bt.Insert(item(10))

	bt.Print(os.Stdout)
	// Output: NODE:[4]
	//   NODE:[2]
	//     NODE:[1]
	//     NODE:[3]
	//   NODE:[6 8]
	//     NODE:[5]
	//     NODE:[7]
	//     NODE:[9 10]
}

func ExampleDeleteLeaf() {
	bt := btree.New(2)
	bt.Insert(item(1))
	bt.Insert(item(2))
	bt.Insert(item(3))

	bt.Delete(item(2))
	bt.Print(os.Stdout)
	// Output: NODE:[1 3]
}

func ExampleDeleteNode() {
	bt := btree.New(2)
	bt.Insert(item(1))
	bt.Insert(item(2))
	bt.Insert(item(3))
	bt.Insert(item(4))
	bt.Insert(item(5))
	bt.Insert(item(6))
	bt.Insert(item(7))
	bt.Insert(item(8))
	bt.Insert(item(9))
	bt.Insert(item(10))

	bt.Delete(item(8))
	bt.Print(os.Stdout)
	// Output: NODE:[4]
	//   NODE:[2]
	//     NODE:[1]
	//     NODE:[3]
	//   NODE:[6 9]
	//     NODE:[5]
	//     NODE:[7]
	//     NODE:[10]
}
