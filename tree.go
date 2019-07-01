package tree

import "fmt"

//Record is a struct containing int fields ID and Parent
type Record struct {
	ID int64
}

//Node is a struct containing int field ID and []*Node field Children
type Node struct {
	left  *Node
	right *Node
	ID    int64
}

func (n *Node) Insert(rec Record) {
	if n == nil {
		return
	}
	//insert on the left
	if rec.ID <= n.ID {
		if n.left == nil {
			n.left = &Node{ID: rec.ID, left: nil, right: nil}
		} else {
			n.left.Insert(rec)
		}
		//insert on the right
	} else {
		if n.right == nil {
			n.right = &Node{ID: rec.ID, left: nil, right: nil}
		} else {
			n.right.Insert(rec)
		}
	}
}

func (n *Node) Print(direction string) {
	if n == nil {
		return
	}
	fmt.Println(direction, " ID: ", n.ID)
	if n.left != nil {
		n.left.Print("left")
	}
	if n.right != nil {
		n.right.Print("right")
	}
}

type Tree struct {
	Root *Node
}

//Build a tree from an unsorted list of records, or add to the current tree
func (t *Tree) Build(records ...Record) {
	//nothing to do
	if len(records) <= 0 {
		return
	}
	for _, v := range records {
		if t.Root == nil {
			t.Root = &Node{ID: v.ID, left: nil, right: nil}
			continue
		}
		t.Root.Insert(v)

	}
	return
}

func (t *Tree) Print() {

	if t.Root == nil {
		return
	}
	fmt.Println("Root")
	fmt.Println(t.Root.ID)
	t.Root.left.Print("left")
	t.Root.right.Print("right")
}
