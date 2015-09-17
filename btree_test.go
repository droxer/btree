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

	fmt.Printf("%s%s", strings.Repeat(" ", 4), strings.Join(items, "--"))

	if len(n.children) > 0 {
		fmt.Println()
		level++
		for _, v := range n.children {
			printNode(v, level)
		}
	}

}

func ExampleInsert() {
	btree := New(2)
	btree.Insert(item(1))
	btree.Insert(item(2))
	btree.Insert(item(3))
	btree.Insert(item(4))
	btree.Insert(item(5))
	btree.Insert(item(8))
	btree.Insert(item(9))

	print(btree)
	// Output: 2--4
	//     1    3    5--8--9
}
