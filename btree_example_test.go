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

func TestGet(t *testing.T) {
	bt := btree.New(2)
	bt.Insert(item(4))
	bt.Insert(item(6))
	bt.Insert(item(8))
	bt.Insert(item(18))
	bt.Insert(item(20))
	bt.Insert(item(22))
	bt.Insert(item(24))
	bt.Insert(item(26))
	bt.Insert(item(28))
	bt.Insert(item(30))

	if bt.Get(item(11)) != nil {
		t.Fatalf("expected is nil, actual is %v", bt.Get(item(7)))
	}

	if bt.Get(item(8)) != item(8) {
		t.Fatalf("expected is 5, actual is %v", bt.Get(item(7)))
	}
}

func ExampleInsert() {
	bt := btree.New(2)
	bt.Insert(item(4))
	bt.Insert(item(6))
	bt.Insert(item(8))
	bt.Insert(item(18))
	bt.Insert(item(20))
	bt.Insert(item(22))
	bt.Insert(item(24))
	bt.Insert(item(26))
	bt.Insert(item(28))
	bt.Insert(item(30))

	bt.Insert(item(19))
	bt.Insert(item(21))
	bt.Print(os.Stdout)
	// Output: NODE:[18]
	//   NODE:[6]
	//     NODE:[4]
	//     NODE:[8]
	//   NODE:[22 26]
	//     NODE:[19 20 21]
	//     NODE:[24]
	//     NODE:[28 30]
}

func ExampleDelete() {
	bt := btree.New(2)
	bt.Insert(item(2))
	bt.Insert(item(4))
	bt.Insert(item(6))
	bt.Insert(item(8))
	bt.Insert(item(18))
	bt.Insert(item(20))
	bt.Insert(item(22))
	bt.Insert(item(24))
	bt.Insert(item(26))
	bt.Insert(item(28))

	bt.Delete(item(4))
	bt.Print(os.Stdout)
	// Output: NODE:[20]
	//   NODE:[8]
	//     NODE:[2 6]
	//     NODE:[18]
	//   NODE:[24]
	//     NODE:[22]
	//     NODE:[26 28]
}
