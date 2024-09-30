package ds

import "errors"

type ListNode struct {
	Value int
	Next  *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (l *LinkedList) InsertAtEnd(val int) {
	//If list is empty just place new node in first position
	if l.Head == nil {
		l.Head = &ListNode{Value: val, Next: nil}
		return
	}

	//Use a temp to traverse instead of l.Head
	tmp := l.Head

	//Iterate until end
	for tmp.Next != nil {
		tmp = tmp.Next
	}

	//Set end to new node with given val
	tmp.Next = &ListNode{Value: val, Next: nil}
}

func (l *LinkedList) InsertAtStart(val int) {
	//If list is empty just place new node in first position
	if l.Head == nil {
		l.Head = &ListNode{Value: val, Next: nil}
		return
	}

	newHead := &ListNode{Value: val, Next: l.Head}
	l.Head = newHead
}

// Remove the first occurrence of the provided Value
func (l *LinkedList) Delete(val int) error {
	//check if linked list is empty
	if l.Head == nil {
		return errors.New("cannot delete from empty linked list")
	}
	//Check if Head is the Value to be deleted
	if l.Head.Value == val {
		l.Head = l.Head.Next
		return nil
	}

	// for init; condition; increment {...}
	for prev := l.Head; prev.Next != nil; prev = prev.Next {
		if prev.Next.Value == val {
			prev.Next = prev.Next.Next
			return nil
		}
	}

	//return error if Value is not found
	return errors.New("provided Value does not exist in linked list")
}

func (l *LinkedList) Contains(Value int) bool {
	if l.Head == nil {
		return false
	}

	for tmp := l.Head; tmp != nil; tmp = tmp.Next {
		if tmp.Value == Value {
			return true
		}
	}

	return false

}

func (l *LinkedList) IsEmpty() bool {
	return l.Head == nil
}

func (l *LinkedList) Length() int {
	count := 0
	for tmp := l.Head; tmp != nil; tmp = tmp.Next {
		count += 1
	}
	return count
}
