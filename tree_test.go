package binTree

import (
	"testing"
	)

var (
	tree = Tree{
		&Node{nil, 0, nil, nil, 1},
		func(lhs interface{}, rhs interface{}) bool {
			return lhs.(int) < rhs.(int)
		},
	}
)

func TestTree_Init(t *testing.T) {
	if tree.root.Value != 0 {
		t.Error("Expected tree.root.Value", 0, ",got", tree.root.Value)
	}
}

func TestTree_Find(t *testing.T) {
	val := tree.Find(0)
	if val.Value != 0 {
		t.Error("Expected", 0, ",got", val.Value)
	}
}

func TestTree_Insert(t *testing.T) {
	tree.insert(5, tree.root)
	tree.InsertValues(1, 2, 3, 4, 5)
	for i := 0; i < 5; i++ {
		val := tree.Find(i)
		if val.Value != i {
			t.Error("Expected", i, ",got", val.Value)
		}
	}
	val := tree.Find(5)
	if val.Value != 5 {
		t.Error("Expected", 5, ",got", val.Value)
	}
	if val.count != 2 {
		t.Error("Expected", 2, ",got", val.count)
	}
}

func TestTree_Delete(t *testing.T) {
	tree.Delete(0)
	val := tree.Find(0)
	if val != nil {
		t.Error("Expected", nil, ",got", val)
	}
	if tree.root == nil {
		t.Error("Expected tree root", ",got", tree.root)
	}
	tree.Delete(5)
	val = tree.Find(5)
	if val == nil {
		t.Error("Expected", 5, ",got", nil)
	}
	for i := 1; i <= 5; i++ {
		tree.Delete(i)
	}
	for i := 0; i <= 5; i++ {
		val = tree.Find(i)
		if val != nil {
			t.Error("Expected nil node", i, ",got", val)
		}
	}
}

func TestTree_String(t *testing.T) {
	tr := Tree{
		&Node{nil, "a", nil, nil, 1},
		func(lhs interface{}, rhs interface{}) bool {
			return lhs.(string) < rhs.(string)
		},
	}
	tr.InsertValues("b", "c")
	val := tr.Find("a")
	if val.Value != "a" {
		t.Error("Expected", `"a"`, ",got", val.Value)
	}
	val = tr.Find("b")
	if val.Value != "b" {
		t.Error("Expected", `"b"`, ",got", val.Value)
	}
	val = tr.Find("c")
	if val.Value != "c" {
		t.Error("Expected", `"c"`, ",got", val.Value)
	}
}

func TestTree_Struct(t *testing.T) {
	type Pack struct {
		num int
		str string
	}
	tr := Tree{
		&Node{nil, Pack{0, "s"}, nil, nil, 1},
		func(lhs interface{}, rhs interface{}) bool {
			return lhs.(Pack).num < rhs.(Pack).num
		},
	}
	tr.InsertValues(Pack{1, "a"}, Pack{2, "b"})
	val := tr.Find(Pack{1, "a"})
	if val.Value.(Pack).str != "a" {
		t.Error("Expected", `"a"`, ",got", val.Value)
	}
	val = tr.Find(Pack{2, "b"})
	if val.Value.(Pack).str != "b" {
		t.Error("Expected", `"b"`, ",got", val.Value)
	}
	val = tr.Find(Pack{0, "s"})
	if val.Value.(Pack).str != "s" {
		t.Error("Expected", `"s"`, ",got", val.Value)
	}

	tr.Delete(Pack{0, "s"})
	val = tr.Find(Pack{0, "s"})
	if val != nil {
		t.Error("Expected", nil, ",got", val)
	}
	val1 := tr.Find(Pack{1, "a"})
	val2 := tr.Find(Pack{2, "b"})
	if val1 != tr.root && val2 != tr.root {
		t.Error("Expected root, got nothing.")
	}
}

// Problems with string slices.
func TestTree_Types(t *testing.T) {
	t.Skip()
	tr := Tree{
		nil,
		func(lhs interface{}, rhs interface{}) bool {
			return lhs.(string) < rhs.(string)
		},
	}
	args := []string{"a", "b", "c"}
	tr.InsertValues(args)
	for _, el := range args {
		n := tr.FindGreatOrEqual(el)
		if n.Value != el {
			t.Error("Expected", el, ", got", n.Value)
		}
	}
}

func TestTree_FindGE(t *testing.T) {
	tr := Tree{
		nil,
		func(lhs interface{}, rhs interface{}) bool {
			return lhs.(int) < rhs.(int)
		},
	}
	args := []int{-1, 3, 6, 10, 15, 20, 25, 30, 35, 40}
	tr.InsertValues(args)
	for _, num := range args {
		n := tr.FindGreatOrEqual(num)
		if n.Value != num {
			t.Error("Expected", num, ", got", n.Value)
		}
		n = tr.FindGreatOrEqual(num - 2)
		if n.Value != num {
			t.Error("Expected", num, ", got", n.Value)
		}
	}
}
