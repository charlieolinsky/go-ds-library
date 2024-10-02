package ds

import (
	"errors"
)

type TreeNode struct {
	Value    int
	Parent   *TreeNode
	Children []*TreeNode
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) Insert(parent *TreeNode, val int) *TreeNode {
	node := &TreeNode{
		Value:    val,
		Parent:   parent,
		Children: []*TreeNode{},
	}

	//When parent is nil we are looking at the root node
	if parent == nil {
		//If the root is empty then we need a new root
		if t.root == nil {
			t.root = node
		}
	} else {
		//Add a new node to the parents children
		parent.Children = append(parent.Children, node)
	}

	//Return the newly created Node
	return node
}

// Return true or false if the specified value exists in the tree.
// Recursive DFS: Worst-case Time O(n) where n is the number of nodes in the tree. Worse-case Space O(d) where d is the depth of the tree
func (t *Tree) Contains(root *TreeNode, val int) bool {
	//Base case: If the current node is nil, return false
	if root == nil {
		return false
	}

	// If the current node's value is equal to the target value, return true
	if root.Value == val {
		return true
	}

	//Recursively Search the children of the current node
	for i := 0; i < len(t.root.Children); i++ {
		if t.Contains(root.Children[i], val) {
			return true
		}
	}

	//If no match is found in the current node or its children, return false
	return false
}

// Delete the first occurrence of the provided value in the tree or return error
// Recursive DFS: time O(n), worst Case Space O(n), average case space O(log(n))
func (t *Tree) Delete(root *TreeNode, val int) error {
	if root == nil {
		return errors.New("cannot delete from an empty tree")
	}

	// Traverse the children to find the node with the value
	for i := 0; i < len(root.Children); i++ {

		//Current Node (potential node to be deleted)
		cur := root.Children[i]

		if cur.Value == val {
			// take a slice of Children up to but not including cur (i), then append all of cur's children to the root's children slice.
			//This effectively deletes cur and adds its children in its place
			root.Children = append(root.Children[:i], cur.Children...)

			// Update parent pointers of promoted nodes. Once cur is deleted is children are still pointing at it.
			for _, child := range cur.Children {
				child.Parent = root
			}
			return nil
		}

		// Recursively search in cur's children
		if err := t.Delete(cur, val); err == nil {
			return nil
		}
	}

	return errors.New("provided value does not exist in the tree")
}
