package ds

import "errors"

type ListNode struct {
	value int
	next  *ListNode
}

type LinkedList struct {
	head *ListNode
}

func (l *LinkedList) InsertAtEnd(val int) {
	//If list is empty just place new node in first position
	if l.head == nil {
		l.head = &ListNode{value: val, next: nil}
		return
	}

	//Use a temp to traverse instead of l.head
	tmp := l.head

	//Iterate until end
	for tmp.next != nil {
		tmp = tmp.next
	}

	//Set end to new node with given val
	tmp.next = &ListNode{value: val, next: nil}
}

func (l *LinkedList) InsertAtStart(val int) {
	//If list is empty just place new node in first position
	if l.head == nil {
		l.head = &ListNode{value: val, next: nil}
		return
	}

	newHead := &ListNode{value: val, next: l.head}
	l.head = newHead
}

// Remove the first occurrence of the provided value
func (l *LinkedList) Delete(val int) error {
	//check if linked list is empty
	if l.head == nil {
		return errors.New("cannot delete from empty linked list")
	}
	//Check if head is the value to be deleted
	if l.head.value == val {
		l.head = l.head.next
		return nil
	}

	// for init; condition; increment {...}
	for prev := l.head; prev.next != nil; prev = prev.next {
		if prev.next.value == val {
			prev.next = prev.next.next
			return nil
		}
	}

	//return error if value is not found
	return errors.New("provided value does not exist in linked list")
}

func (l *LinkedList) Contains(value int) bool {
	if l.head == nil {
		return false
	}

	for tmp := l.head; tmp != nil; tmp = tmp.next {
		if tmp.value == value {
			return true
		}
	}

	return false

}

func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

func (l *LinkedList) Length() int {
	count := 0
	for tmp := l.head; tmp != nil; tmp = tmp.next {
		count += 1
	}
	return count
}
