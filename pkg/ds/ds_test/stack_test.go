package ds

import (
	"testing"

	"github.com/charlieolinsky/go-ds-library/pkg/ds"
)

func TestStackPush(t *testing.T) {
	stack := ds.Stack{}
	stack.Push(1)

	if len(stack.Data) != 1 {
		t.Errorf("Expected stack length to be 1 after pushing an item, got %d", len(stack.Data))
	}
}

func TestStackPop(t *testing.T) {
	stack := ds.Stack{}
	stack.Push(1)

	popped, err := stack.Pop()

	if err != nil {
		t.Error("Expected no error when popping from a non-empty stack")
	}
	if popped != 1 {
		t.Errorf("Expected popped value to be 1, got %d", popped)
	}
	if len(stack.Data) != 0 {
		t.Errorf("Expected stack length to be 0 after popping an item, got %d", len(stack.Data))
	}
}

func TestStackPeek(t *testing.T) {
	stack := ds.Stack{}
	stack.Push(1)

	peeked, err := stack.Peek()

	if err != nil {
		t.Error("Expected no error when peeking from a non-empty stack")
	}
	if peeked == 0 {
		t.Errorf("Expected peeked value to be 1, got %d", peeked)
	}
	if len(stack.Data) != 1 {
		t.Errorf("Expected stack length to be 1 after peeking, got %d", len(stack.Data))
	}
}

func TestStackIsEmpty(t *testing.T) {
	stack := ds.Stack{}

	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty")
	}

	stack.Push(1)

	if stack.IsEmpty() == true {
		t.Error("Expected stack not to be empty after pushing an item")
	}
}
