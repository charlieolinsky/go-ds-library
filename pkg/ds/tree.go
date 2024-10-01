package ds

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

// Return true or false if the specified value exists in the true.
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
