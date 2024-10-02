package ds

import "errors"

type Queue[T any] struct {
	Data []T
}

func (q *Queue[T]) Enqueue(x T) {
	q.Data = append(q.Data, x)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.Data) == 0 {
		return *new(T), errors.New("cannot dequeue from empty queue")
	}

	ret := q.Data[0]
	q.Data = q.Data[1:]

	return ret, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if len(q.Data) == 0 {
		return *new(T), errors.New("cannot peek from an empty queue")
	}

	return q.Data[0], nil

}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.Data) == 0
}
