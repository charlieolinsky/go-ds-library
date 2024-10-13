package ds

import (
	"testing"

	"github.com/charlieolinsky/go-ds-library/pkg/ds"
)

func TestInsertAndContains(t *testing.T) {
	tree := ds.Tree{}

	// Insert Root
	tree.Insert(nil, 1)

	// Test contains
	if !tree.Contains(tree.Root, 1) {
		t.Errorf("Expected tree to contain value 1")
	}

	// Insert children
	child1 := tree.Insert(tree.Root, 2)
	tree.Insert(tree.Root, 3)

	// Test contains for children
	if !tree.Contains(tree.Root, 2) {
		t.Errorf("Expected tree to contain value 2")
	}
	if !tree.Contains(tree.Root, 3) {
		t.Errorf("Expected tree to contain value 3")
	}

	// Test contains for a non-existent value
	if tree.Contains(tree.Root, 4) {
		t.Errorf("Did not expect tree to contain value 4")
	}

	// Insert more children
	tree.Insert(child1, 4)
	tree.Insert(child1, 5)

	// Test deeper levels
	if !tree.Contains(tree.Root, 4) {
		t.Errorf("Expected tree to contain value 4")
	}
	if !tree.Contains(tree.Root, 5) {
		t.Errorf("Expected tree to contain value 5")
	}

	// Test for duplicate insert
	tree.Insert(child1, 4) // Assuming duplicates are allowed
	if !tree.Contains(tree.Root, 4) {
		t.Errorf("Expected tree to still contain value 4 after duplicate insert")
	}
}

func TestTreeDelete(t *testing.T) {
	tree := ds.Tree{}
	tree.Insert(nil, 1)
	child1 := tree.Insert(tree.Root, 2)
	tree.Insert(child1, 4)
	tree.Insert(child1, 5)

	// Delete a child node
	if err := tree.Delete(tree.Root, 2); err != nil {
		t.Errorf("Expected to delete node with value 2: %v", err)
	}

	if tree.Contains(tree.Root, 2) {
		t.Errorf("Expected tree not to contain value 2 after deletion")
	}

	// Check if children of deleted node are still present
	if !tree.Contains(tree.Root, 4) || !tree.Contains(tree.Root, 5) {
		t.Errorf("Expected tree to still contain children of deleted node")
	}

	// Delete the root node
	if err := tree.Delete(tree.Root, 1); err != nil {
		t.Errorf("Expected to delete root node with value 1: %v", err)
	}

	if tree.Contains(tree.Root, 1) {
		t.Errorf("Expected tree not to contain value 1 after deletion")
	}

	// Attempt to delete a non-existent value
	if err := tree.Delete(tree.Root, 10); err == nil {
		t.Errorf("Expected an error when trying to delete a non-existent value")
	}
}

func TestEmptyTreeDelete(t *testing.T) {
	tree := ds.Tree{}

	// Attempt to delete from an empty tree
	if err := tree.Delete(tree.Root, 1); err == nil {
		t.Errorf("Expected an error when trying to delete from an empty tree")
	}
}
