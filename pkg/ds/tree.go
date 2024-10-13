package ds

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	Value    int
	Parent   *TreeNode
	Children []*TreeNode
}

type Tree struct {
	Root *TreeNode
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
		if t.Root == nil {
			t.Root = node
		}
	} else {
		//Add a new node to the parents children
		parent.Children = append(parent.Children, node)
	}

	//Return the newly created Node
	return node
}

// Return true or false if the specified value exists in the tree.
// Recursive DFS: worst-case time O(n). worst-case space O(d) where d is the depth of the tree
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
	for i := 0; i < len(root.Children); i++ {
		if t.Contains(root.Children[i], val) {
			return true
		}
	}

	//If no match is found in the current node or its children, return false
	return false
}

// Delete the first occurrence of the provided value in the tree or return error
// Recursive DFS: time O(n), worst-case space O(n), average-case space O(log(n))
func (t *Tree) Delete(root *TreeNode, val int) error {
	if t.Root == nil {
		return errors.New("cannot delete from an empty tree")
	}

	// Check for case where the root is to be deleted
	if t.Root.Value == val {
		if len(t.Root.Children) > 0 {
			// Promote the first child as the new root
			newRoot := t.Root.Children[0]
			newRoot.Parent = nil // Set new root's parent to nil

			// Append the rest of the old root's children to the new root
			newRoot.Children = append(newRoot.Children, t.Root.Children[1:]...)

			t.Root = newRoot // Update the tree's root
		} else {
			t.Root = nil // Tree is now empty
		}
		return nil
	}

	// Recursive helper function for deletion when root is not val
	return t.deleteNode(root, val)
}

// Helper function to delete a node
func (t *Tree) deleteNode(root *TreeNode, val int) error {
	for i := 0; i < len(root.Children); i++ {
		cur := root.Children[i]

		if cur.Value == val {
			// Remove cur and append its children
			root.Children = append(root.Children[:i], append(cur.Children, root.Children[i+1:]...)...)
			// Update parent pointers
			for _, child := range cur.Children {
				child.Parent = root
			}
			return nil
		}

		// Recursively search in cur's children
		if err := t.deleteNode(cur, val); err == nil {
			return nil
		}
	}

	return errors.New("provided value does not exist in the tree")
}

func (t *Tree) PreOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("PreO: %d\n", root.Value)

	for _, child := range root.Children {
		t.PreOrderTraversal(child)
	}
}

func (t *Tree) PostOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	for _, child := range root.Children {
		t.PostOrderTraversal(child)
	}
	fmt.Printf("PostO: %d\n", root.Value)
}

func (t *Tree) LevelOrderTraversal() {
	if t.Root == nil {
		return
	}

	q := Queue[*TreeNode]{}
	q.Enqueue(t.Root)

	for !q.IsEmpty() {
		cur, _ := q.Dequeue()

		fmt.Printf("LO: %d\n", cur.Value)

		for _, child := range cur.Children {
			q.Enqueue(child)
		}
	}
}

func (t *Tree) Height(root *TreeNode) int {
	if root == nil {
		return -1
	}
	if len(root.Children) == 0 {
		return 0
	}

	maxHeight := -1

	for _, child := range root.Children {
		childHeight := t.Height(child)
		maxHeight = max(childHeight, maxHeight)
	}
	return maxHeight + 1

}
