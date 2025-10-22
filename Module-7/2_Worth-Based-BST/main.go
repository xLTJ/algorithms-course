package main

import "fmt"

type FamilyNode struct {
	Key             string
	Worth           int
	CumulativeWorth int
	Size            int
	Parent          *FamilyNode
	Left            *FamilyNode
	Right           *FamilyNode
}

type FamilyBST struct {
	Root *FamilyNode
}

// NewFamilyNode creates and returns a pointer to a new FamilyNode with the specified key and worth.
// Why? Cus im lazy (also safer code yay we love that)
func NewFamilyNode(key string, worth int) *FamilyNode {
	return &FamilyNode{Key: key, Worth: worth}
}

// Insert adds a new FamilyNode to the FamilyBST. Keeps the BST property thing intact.
// Also updates the CumulativeWorth and Size properties along the insertion path.
func (t *FamilyBST) Insert(node *FamilyNode) {
	var y *FamilyNode = nil
	x := t.Root

	// find spot for node
	for x != nil {
		y = x
		if node.Worth < x.Worth {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	// insert node
	node.Parent = y
	if y == nil {
		t.Root = node
	} else if node.Worth < y.Worth {
		y.Left = node
	} else {
		y.Right = node
	}

	// update cumulative networth on path
	node.CumulativeWorth = node.Worth
	x = node
	for x.Parent != nil {
		x.Parent.CumulativeWorth += node.Worth
		x.Parent.Size++
		x = x.Parent
	}
}

// Min returns the node with the minimum worth in the subtree rooted at the provided node.
// Use t.Root as the node if you want overall min
func (t *FamilyBST) Min(x *FamilyNode) *FamilyNode {
	for x.Left != nil {
		x = x.Left
	}
	return x
}

// Max returns the node with the maximum worth in the subtree rooted at the provided node.
// Use t.Root as the node if you want overall max
func (t *FamilyBST) Max(x *FamilyNode) *FamilyNode {
	for x.Right != nil {
		x = x.Right
	}
	return x
}

// Rank calculates the rank of the given node in the tree based on its worth.
// Lower rank = higher worth
// Returns an integer representing the rank or an error if the node is not found.
func (t *FamilyBST) Rank(node *FamilyNode) (int, error) {
	rank := 0
	current := t.Root

	for current != nil {
		if node.Key == current.Key {
			rightSize := 0
			if current.Right != nil {
				rightSize = current.Right.Size
			}

			return rank + rightSize + 1, nil
		} else if node.Worth > current.Worth {
			current = current.Right
		} else {
			rightSize := 0
			if current.Right != nil {
				rightSize = current.Right.Size
			}

			rank += rightSize + 1
			current = current.Left
		}
	}

	return -1, fmt.Errorf("node with this key could not be found")
}

// transplant replaces the subtree rooted at node `u` with the subtree rooted at node `v` in the binary search tree.
// it updates parent-child relationships to maintain the tree structure.
// go look in the book or something for better explanation.
// this is a helper function used for Delete, idk why u would want to use this by itself
func (t *FamilyBST) transplant(u, v *FamilyNode) {
	if u.Parent == nil {
		t.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}

	if v != nil {
		v.Parent = u.Parent
	}
}

// Delete kills the `z` node while maintaining the BST property stuff
func (t *FamilyBST) Delete(z *FamilyNode) {
	if z.Left == nil {
		t.transplant(z, z.Right)
	} else if z.Right == nil {
		t.transplant(z, z.Left)
	} else {
		y := t.Min(z.Right)
		if y.Parent != z {
			t.transplant(y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}

		t.transplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
	}
}

// Traverse is just used for testing, it prints out shit i guess
func (t *FamilyBST) Traverse(node *FamilyNode) {
	if node != nil {
		t.Traverse(node.Left)
		fmt.Println("\nKey:", node.Key, "\nWorth:", node.Worth, "\nC:", node.CumulativeWorth)
		t.Traverse(node.Right)
	}
}

func main() {
	familyMemberNames := []string{"Jan", "Abel", "Rob", "John", "Mary", "Lisa", "Sara", "Sina", "Søren"}
	familyNetWorths := []int{50000, 30000, 20000, 15000, 10000, 5000, 8000, 25000, 7000}
	familyTree := &FamilyBST{}

	for i, name := range familyMemberNames {
		familyTree.Insert(NewFamilyNode(name, familyNetWorths[i]))
	}

	familyTree.Insert(NewFamilyNode("Dan", 12000))

	familyTree.Delete(familyTree.Root.Left.Left)

	fmt.Println("lol")
}
