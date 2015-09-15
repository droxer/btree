package btree

import (
	"sort"
)

// Item represents a single object in the tree.
type Item interface {
	// check wether the current item is less then given param
	Less(other Item) bool
}

type items []Item

func (s items) find(item Item) (index int, found bool) {
	i := sort.Search(len(s), func(i int) bool {
		return item.Less(s[i])
	})

	if i > 0 && !s[i-1].Less(item) {
		return i, true
	}

	return i, false
}

// Node represents single node in BTree.
type Node struct {
	children []*Node
	items    items
	t        *BTree
}

func (n *Node) insert(item Item) {
	i, found := n.items.find(item)
	if found {
		n.items[i] = item
		return
	}

}

type BTree struct {
	degree int
	length int
	root   *Node
}

func New(degree int) *BTree {
	if degree <= 1 {
		panic("invalid degree")
	}

	return &BTree{
		degree: degree,
	}
}

func (b *BTree) min() int {
	return b.degree - 1
}

func (b *BTree) max() int {
	return b.degree*2 - 1
}

func (b *BTree) Insert(item Item) bool {
	if item == nil {
		panic("it not allowed to add nil item.")
	}

	if b.root == nil {
		b.root = &Node{t: b}
		b.root.items = append(b.root.items, item)
		b.length++
		return true
	}

	if len(b.root.items) < b.max() {
		b.root.insert(item)
		return true
	}

	return true
}
