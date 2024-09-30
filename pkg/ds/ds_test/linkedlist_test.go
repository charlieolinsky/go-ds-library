package ds

import (
	"testing"

	"github.com/charlieolinsky/go-ds-library/pkg/ds"
)

func TestInsertAtEnd(t *testing.T) {
	l := ds.LinkedList{}

	//Test insert into empty list
	l.InsertAtEnd(1)
	if l.Head.Value != 1 {
		t.Fatalf("expected head of linked list to be 1, got %d", l.Head.Value)
	}

	//Test insert into a non-empty list
	l.InsertAtEnd(2)
	l.InsertAtEnd(3)

	//Iterate through the linked list and ensure values are inserted in the correct order
	expectedValues := []int{1, 2, 3}
	checkExpected(&l, expectedValues, t)
}

func TestInsertAtStart(t *testing.T) {
	l := ds.LinkedList{}

	//Check insert from empty
	l.InsertAtStart(1)
	if l.Head.Value != 1 {
		t.Errorf("expected 1 inserted at start of list, got %d", l.Head.Value)
	}

	//Check insert from non-empty and ensure proper order
	l.InsertAtStart(2)
	l.InsertAtStart(3)

	expectedValues := []int{3, 2, 1}
	checkExpected(&l, expectedValues, t)
}

func TestDelete(t *testing.T) {
	l := ds.LinkedList{}

	//Test delete from empty list
	if err := l.Delete(1); err == nil {
		t.Errorf("expected error when deleting from an empty linked list")
	}

	//Test deleting head
	l.InsertAtEnd(1)
	l.InsertAtEnd(2)

	// 1 => 2 => nil

	if _ = l.Delete(1); l.Head.Value != 2 {
		t.Errorf("expected list head to equal 2 after deletion, got %d", l.Head.Value)
	}

	// 2 => nil

	//Test deletion from middle
	l.InsertAtEnd(3)
	l.InsertAtEnd(4)

	//2 => 3 => 4 => nil

	_ = l.Delete(3)
	expectedValues := []int{2, 4}
	checkExpected(&l, expectedValues, t)

	//2 => 4 => nil

	//Test deletion from end
	_ = l.Delete(4)
	expectedValues = []int{2}
	checkExpected(&l, expectedValues, t)

	//2 => nil

}

// Example of table-driven testing
func TestContains(t *testing.T) {
	l := ds.LinkedList{}

	cases := []struct {
		insertValues []int
		searchValue  int
		expected     bool
	}{
		{[]int{}, 1, false},
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, 5, false},
	}

	for _, c := range cases {
		l = ds.LinkedList{}

		for _, val := range c.insertValues {
			l.InsertAtEnd(val)
		}

		found := l.Contains(c.searchValue)
		if found != c.expected {
			t.Errorf("expected Contains(%d) to return %t, got %t", c.searchValue, c.expected, found)
		}
	}

}

func TestIsEmpty(t *testing.T) {
	l := ds.LinkedList{}

	if l.IsEmpty() != true {
		t.Errorf("expected true when checking if an empty list is empty, got %t", l.IsEmpty())
	}

	l.InsertAtEnd(1)

	if l.IsEmpty() != false {
		t.Errorf("expected false when checking if a non-empty list is empty, got %t", l.IsEmpty())
	}
}

func TestLength(t *testing.T) {
	l := ds.LinkedList{}

	if l.Length() != 0 {
		t.Errorf("expected length 0 when checking length of empty list, got %d", l.Length())
	}

	l.InsertAtEnd(1)

	if l.Length() != 1 {
		t.Errorf("expected length 1 when checking length of non-empty list, got %d", l.Length())
	}
}

// helper method, checks if values in the list match those that are expected.
func checkExpected(l *ds.LinkedList, expectedValues []int, t *testing.T) {
	// Check if the length of the linked list matches the expected length
	count := 0
	for tmp := l.Head; tmp != nil; tmp = tmp.Next {
		count++
	}

	if count != len(expectedValues) {
		t.Errorf("expected list length to be %d, got %d", len(expectedValues), count)
	}

	// Iterate through the linked list and check values
	i := 0
	for tmp := l.Head; tmp != nil; tmp = tmp.Next {
		if i >= len(expectedValues) {
			t.Errorf("unexpected additional node with value %d at position %d", tmp.Value, i)
			break
		}

		if tmp.Value != expectedValues[i] {
			t.Errorf("expected value at position %d to equal %d, got %d", i, expectedValues[i], tmp.Value)
		}
		i++
	}
}
