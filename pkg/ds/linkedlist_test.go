package ds

import "testing"

func TestInsertAtEnd(t *testing.T) {
	l := LinkedList{}

	//Test insert into empty list
	l.InsertAtEnd(1)
	if l.head.value != 1 {
		t.Errorf("expected head of linked list to be 1, got %d", l.head.value)
	}

	//Test insert into a non-empty list
	l.InsertAtEnd(2)
	l.InsertAtEnd(3)
	l.InsertAtEnd(4)

	//Iterate through the linked list and ensure values are inserted in the correct order
	expectedValues := []int{1, 2, 3, 4}
	i := 0

	for tmp := l.head; tmp != nil; tmp = tmp.next {
		if tmp.value != expectedValues[i] {
			t.Errorf("expected node value at position %d to be %d, got %d", i, expectedValues[i], tmp.value)
		}
		i++
	}

	//Ensure the list has exactly 4 elements
	if i != len(expectedValues) {
		t.Errorf("expected list length to be %d, got %d", len(expectedValues), i)
	}

}
