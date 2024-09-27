package ds

import "errors"

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(x int) {
	q.data = append(q.data, x)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.data) == 0 {
		return 0, errors.New("cannot dequeue from empty queue")
	}

	ret := q.data[0]
	q.data = q.data[1:]

	return ret, nil
}

func (q *Queue) Peek() (int, error) {
	if len(q.data) == 0 {
		return -1, errors.New("cannot peek from an empty queue")
	}

	return q.data[0], nil

}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
