package btree

import (
	"fmt"
	"strings"
)

type item int

func (i item) Less(other Item) bool {
	return i < other.(item)
}

func print(b *BTree) {
	printNode(b.root, 0)
}

func printNode(n *node, level int) {
	var items []string
	for _, v := range n.items {
		items = append(items, fmt.Sprintf("%v", v))
	}

	fmt.Printf("level %d: %s\n", level, strings.Join(items, "--"))

	level++
	for _, v := range n.children {
		printNode(v, level)
	}
}

func ExamplePrint() {
	btree := New(2)
	btree.Insert(item(1))
	btree.Insert(item(2))
	btree.Insert(item(3))
	btree.Insert(item(4))
	btree.Insert(item(5))
	btree.Insert(item(8))
	btree.Insert(item(9))

	print(btree)
	// Output: level 0: 2--4
	// level 1: 1
	// level 1: 3
	// level 1: 5--8--9
}
