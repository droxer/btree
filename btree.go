package btree

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

// Item represents a single object in the tree.
type Item interface {
	Less(other Item) bool
}

type items []Item

func (s items) find(item Item) (index int, found bool) {
	i := sort.Search(len(s), func(i int) bool {
		return item.Less(s[i])
	})

	if i > 0 && !s[i-1].Less(item) {
		return i - 1, true
	}

	return i, false
}

func (s *items) insertAt(index int, item Item) {
	*s = append(*s, nil)

	if index < len(*s) {
		copy((*s)[index+1:], (*s)[index:])
	}
	(*s)[index] = item
}

func (s *items) removeAt(index int) Item {
	out := (*s)[index]
	copy((*s)[index:], (*s)[index+1:])
	*s = (*s)[:len(*s)-1]
	return out
}

func (s *items) pop() (out Item) {
	last := len(*s) - 1
	out, *s = (*s)[last], (*s)[:last]
	return
}

type children []*node

func (s *children) insertAt(i int, n *node) {
	*s = append(*s, nil)

	if i < len(*s) {
		copy((*s)[i+1:], (*s)[i:])
	}
	(*s)[i] = n
}

func (s *children) removeAt(i int) *node {
	out := (*s)[i]
	copy((*s)[i:], (*s)[i+1:])
	(*s) = (*s)[:len(*s)-1]
	return out
}

func (s *children) pop() (out *node) {
	last := len(*s) - 1
	out, *s = (*s)[last], (*s)[:last]
	return
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

func (n *node) remove(key Item) Item {
	i, found := n.items.find(key)

	if len(n.children) == 0 {
		if found {
			return n.items.removeAt(i)
		}
		return nil
	}

	child := n.children[i]
	if len(child.items) <= n.t.min() {
		n.growAndRemove(i, key)
	}

	if found {
		return n.items[i]
	}

	return n.children[i].remove(key)
}

func (n *node) growAndRemove(i int, item Item) Item {
	child := n.children[i]

	if len(n.items) > i && len(n.children[i+1].items) > n.t.min() {
		sep := n.items[i]
		right := n.children[i+1]

		child.items = append(child.items, sep)
		n.items[i] = right.items.removeAt(0)

		if len(n.children[i+1].children) > 0 {
			child.children = append(child.children, right.children.removeAt(0))
		}

	} else if i > 0 && len(n.children[i-1].items) > n.t.min() {
		sep := n.items[i-1]
		left := n.children[i-1]

		child.items.insertAt(0, sep)
		n.items[i-1] = left.items.pop()

		if len(n.children[i-1].children) > 0 {
			child.children.insertAt(0, left.children.pop())
		}

	} else {
		sep := n.items.removeAt(i)
		right := n.children.removeAt(i + 1)

		child.items = append(child.items, sep)
		child.items = append(child.items, right.items...)
		child.children = append(child.children, right.children...)
	}

	return child.remove(item)
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

func (n *node) get(key Item) Item {
	i, found := n.items.find(key)
	if found {
		return n.items[i]
	} else if len(n.children) > 0 {
		return n.children[i].get(key)
	}
	return nil
}

func (b *BTree) deleteItem(key Item) Item {
	if b.root == nil || len(b.root.items) == 0 {
		return nil
	}

	out := b.root.remove(key)

	if len(b.root.children) == 0 && len(b.root.items) == 0 {
		b.root = b.root.children[0]
	}

	return out
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

func (b *BTree) Get(key Item) Item {
	if b.root == nil {
		return nil
	}

	return b.root.get(key)
}

func (b *BTree) Delete(key Item) Item {
	return b.deleteItem(key)
}

func (b *BTree) Print(w io.Writer) {
	b.root.print(w, 0)
}

func (n *node) print(w io.Writer, level int) {
	fmt.Fprintf(w, "%sNODE:%v\n", strings.Repeat("  ", level), n.items)
	for _, c := range n.children {
		c.print(w, level+1)
	}
}
