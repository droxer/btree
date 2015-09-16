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

func (s *items) insertAt(i int, item Item) {
	*s = append(*s, nil)

	if i < len(*s) {
		copy((*s)[i+1:], (*s)[i:])
	}
	(*s)[i] = item
}

type children []*node

func (s *children) insertAt(i int, n *node) {
	*s = append(*s, nil)

	if i < len(*s) {
		copy((*s)[i+1:], (*s)[i:])
	}
	(*s)[i] = n
}

// node represents single node in BTree.
type node struct {
	children children
	items    items
	t        *BTree
}

func (n *node) split(i int) (Item, *node) {
	item := n.items[i]
	next := n.t.newNode()
	next.items = append(next.items, n.items[i+1:]...)
	n.items = n.items[:i]

	if len(n.children) > 0 {
		next.children = append(next.children, n.children[i+1:]...)
		n.children = n.children[:i+1]
	}

	return item, next
}

func (n *node) maybeSplit(i int) bool {
	if len(n.children[i].items) < n.t.max() {
		return false
	}

	first := n.children[i]
	item, second := first.split(n.t.max() / 2)
	n.items.insertAt(i, item)
	n.children.insertAt(i+1, second)
	return true
}

func (n *node) insert(item Item) {
	i, found := n.items.find(item)
	if found {
		n.items[i] = item
		return
	}

	if len(n.children) == 0 {
		n.items.insertAt(i, item)
		return
	}

	if n.maybeSplit(i) {
		inItem := n.items[i]
		switch {
		case item.Less(inItem):
		case inItem.Less(item):
			i++
		default:
			n.items[i] = item
			return
		}
	}

	n.children[i].insert(item)
}

type BTree struct {
	degree int
	root   *node
}

func New(degree int) *BTree {
	if degree <= 1 {
		panic("invalid degree")
	}

	return &BTree{
		degree: degree,
	}
}

func (b *BTree) newNode() *node {
	return &node{
		t: b,
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
		b.root = b.newNode()
		b.root.items = append(b.root.items, item)
		return true
	} else if len(b.root.items) >= b.max() {
		newItem, second := b.root.split(b.max() / 2)
		oldRoot := b.root
		b.root = b.newNode()
		b.root.items = append(b.root.items, newItem)
		b.root.children = append(b.root.children, oldRoot, second)
	}

	b.root.insert(item)
	return true
}
