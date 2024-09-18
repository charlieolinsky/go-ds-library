package ds

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := Stack{}
	stack.Push(1)

	if len(stack.data) != 1 {
		t.Errorf("Expected stack length to be 1 after pushing an item, got %d", len(stack.data))
	}
}

func TestPop(t *testing.T) {
	stack := Stack{}
	stack.Push(1)

	popped, err := stack.Pop()

	if err != nil {
		t.Error("Expected no error when popping from a non-empty stack")
	}
	if popped != 1 {
		t.Errorf("Expected popped value to be 1, got %d", popped)
	}
	if len(stack.data) != 0 {
		t.Errorf("Expected stack length to be 0 after popping an item, got %d", len(stack.data))
	}
}

func TestPeek(t *testing.T) {
	stack := Stack{}
	stack.Push(1)

	peeked, err := stack.Peek()

	if err != nil {
		t.Error("Expected no error when peeking from a non-empty stack")
	}
	if peeked == 0 {
		t.Errorf("Expected peeked value to be 1, got %d", peeked)
	}
	if len(stack.data) != 1 {
		t.Errorf("Expected stack length to be 1 after peeking, got %d", len(stack.data))
	}
}

func TestIsEmpty(t *testing.T) {
	stack := Stack{}

	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty")
	}

	stack.Push(1)

	if stack.IsEmpty() == true {
		t.Error("Expected stack not to be empty after pushing an item")
	}
}
