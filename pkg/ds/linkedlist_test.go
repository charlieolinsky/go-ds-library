package ds

import "testing"

func TestInsertAtEnd(t *testing.T) {
	l := LinkedList{}

	//Test insert into empty list
	l.InsertAtEnd(1)
	if l.head.value != 1 {
		t.Fatalf("expected head of linked list to be 1, got %d", l.head.value)
	}

	//Test insert into a non-empty list
	l.InsertAtEnd(2)
	l.InsertAtEnd(3)

	//Iterate through the linked list and ensure values are inserted in the correct order
	expectedValues := []int{1, 2, 3}
	checkExpected(&l, expectedValues, t)
}

func TestInsertAtStart(t *testing.T) {
	l := LinkedList{}

	//Check insert from empty
	l.InsertAtStart(1)
	if l.head.value != 1 {
		t.Errorf("expected 1 inserted at start of list, got %d", l.head.value)
	}

	//Check insert from non-empty and ensure proper order
	l.InsertAtStart(2)
	l.InsertAtStart(3)

	expectedValues := []int{3, 2, 1}
	checkExpected(&l, expectedValues, t)
}

func TestDelete(t *testing.T) {
	l := LinkedList{}

	//Test delete from empty list
	err := l.Delete(1)
	if err == nil {
		t.Errorf("expected error when deleting from an empty linked list")
	}

	//Test deleting head
	l.InsertAtEnd(1)
	l.InsertAtEnd(2)

	_ = l.Delete(1)

	if l.head.value != 2 {
		t.Errorf("expected list head to equal 2 after deletion, got %d", l.head.value)
	}

	//Test deletion from middle
	l.InsertAtEnd(3)
	l.InsertAtEnd(4)

	_ = l.Delete(3)
	expectedValues := []int{2, 4}
	checkExpected(&l, expectedValues, t)
}

// LinkedList helper method, checks if values in the list match those that are expected.
func checkExpected(l *LinkedList, expectedValues []int, t *testing.T) {
	// Check if the length of the linked list matches the expected length
	count := 0
	for tmp := l.head; tmp != nil; tmp = tmp.next {
		count++
	}

	if count != len(expectedValues) {
		t.Errorf("expected list length to be %d, got %d", len(expectedValues), count)
	}

	// Iterate through the linked list and check values
	i := 0
	for tmp := l.head; tmp != nil; tmp = tmp.next {
		if i >= len(expectedValues) {
			t.Errorf("unexpected additional node with value %d at position %d", tmp.value, i)
			break
		}

		if tmp.value != expectedValues[i] {
			t.Errorf("expected value at position %d to equal %d, got %d", i, expectedValues[i], tmp.value)
		}
		i++
	}
}
