package ds

import (
	"testing"

	"github.com/charlieolinsky/go-ds-library/pkg/ds"
)

func TestQueueEnqueue(t *testing.T) {
	q := ds.Queue[int]{}

	q.Enqueue(1)

	if len(q.Data) != 1 {
		t.Errorf("expected queue length to be 1 after enqueue, got %d", len(q.Data))
	} else if q.Data[0] != 1 {
		t.Errorf("expected enqueued value to be 1, got %d", q.Data[0])
	}

}

func TestQueueDequeue(t *testing.T) {
	q := ds.Queue[int]{}

	_, err := q.Dequeue()

	if err == nil {
		t.Errorf("expected error when dequeuing from an empty queue")
	}

	q.Enqueue(1)

	n, err := q.Dequeue()

	if err != nil {
		t.Errorf("expected no error when dequeuing from a non-empty queue")
	}
	if n != 1 {
		t.Errorf("expected dequeued value to be 1, got %d", n)
	}
	if len(q.Data) != 0 {
		t.Errorf("expected queue length to be 0 after dequeueing an item, got %d", len(q.Data))
	}

}

func TestQueuePeek(t *testing.T) {
	q := ds.Queue[int]{}

	_, err := q.Peek()

	if err == nil {
		t.Errorf("expected error when peeking an empty queue")
	}

	q.Enqueue(1)

	n, err := q.Peek()

	if err != nil {
		t.Errorf("expected no error when peeking non-empty queue")
	}
	if n != 1 {
		t.Errorf("expected value to be 1 after peeking got, %d", n)
	}
}

func TestQueueIsEmpty(t *testing.T) {
	q := ds.Queue[int]{}

	if q.IsEmpty() != true {
		t.Errorf("expected queue to be empty")
	}

	q.Enqueue(1)

	if q.IsEmpty() != false {
		t.Errorf("expected queue to have one element")
	}

}
