package binTree

import (
	"reflect"
)

type Node struct {
	left   *Node
	Value  interface{}
	right  *Node
	parent *Node
	count  int
}

// Tree is a binary tree with Interface values and comparator.
type Tree struct {
	root *Node
	comp func(lhs interface{}, rhs interface{}) bool
}

func (t *Tree) Find(x interface{}) *Node {
	n := t.find(x, t.root)
	if n != nil && !t.comp(x, n.Value) && !t.comp(n.Value, x) {
		return n
	}
	return nil
}

func (t *Tree) find(x interface{}, n *Node) *Node {
	if n != nil {
		if t.comp(x, n.Value) {
			if n.left != nil {
				return t.find(x, n.left)
			}
		} else if t.comp(n.Value, x) {
			if n.right != nil {
				return t.find(x, n.right)
			}
		}
	}
	return n
}

func (t *Tree) FindGreatOrEqual(x interface{}) *Node {
	return t.findGE(x, t.root)
}

func (t *Tree) findGE(x interface{}, n *Node) *Node {
	if n != nil {
		if t.comp(n.Value, x) {
			return t.findGE(x, n.right)
		} else if !t.comp(x, n.Value) {
			return n
		}
		if n.left != nil && !t.comp(n.left, x) {
			return t.findGE(x, n.left)
		}
	}
	return n
}

func (t *Tree) insert(x interface{}, n *Node) {
	if t.comp(x, n.Value) {
		if n.left != nil {
			t.insert(x, n.left)
		} else {
			n.left = &Node{nil, x, nil, n, 1}
		}
	} else if t.comp(n.Value, x) {
		if n.right != nil {
			t.insert(x, n.right)
		} else {
			n.right = &Node{nil, x, nil, n, 1}
		}
	} else {
		n.count++
	}
}

// Works only with []int slices.
func (t *Tree) InsertValues(args ...interface{}) {
	if reflect.TypeOf(args[0]).Kind() == reflect.Slice {
		if t.root == nil {
			t.root = &Node{nil, args[0].([]int)[0], nil, nil, 1}
		}
		for _, el := range args[0].([]int) {
			t.insert(el, t.root)
		}
	} else {
		if t.root == nil {
			t.root = &Node{nil, args[0], nil, nil, 1}
		}
		for _, el := range args {
			t.insert(el, t.root)
		}
	}
}

func (t *Tree) Delete(x interface{}) {
	n := t.Find(x)
	if n == nil {
		return
	}
	if n.count > 1 {
		n.count--
	} else {
		t.delete(n)
	}
}

func (t *Tree) delete(n *Node) {
	if n.left == nil {
		if n.right == nil {
			t.delete0(n)
		} else {
			t.delete1(n, n.right)
		}
	} else if n.right != nil {
		t.delete2(n)
	} else {
		t.delete1(n, n.left)
	}
}

func (t *Tree) delete0(n *Node) {
	if n.parent != nil {
		if n.parent.left == n {
			n.parent.left = nil
		} else {
			n.parent.right = nil
		}
	} else {
		t.root = nil
	}
}

func (t *Tree) delete1(n *Node, c *Node) {
	if n.parent != nil {
		if n.parent.left == n {
			n.parent.left = c
			c.parent = n.parent
		} else {
			n.parent.right = c
			c.parent = n.parent
		}
	} else {
		t.root = c
		c.parent = nil
	}
}

func (t *Tree) delete2(n *Node) {
	c := n.right
	for c.left != nil {
		c = c.left
	}
	n.Value = c.Value
	n.count = c.count
	t.delete(c)
}
