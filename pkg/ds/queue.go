package ds

import "errors"

type Queue struct {
	Data []int
}

func (q *Queue) Enqueue(x int) {
	q.Data = append(q.Data, x)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.Data) == 0 {
		return 0, errors.New("cannot dequeue from empty queue")
	}

	ret := q.Data[0]
	q.Data = q.Data[1:]

	return ret, nil
}

func (q *Queue) Peek() (int, error) {
	if len(q.Data) == 0 {
		return -1, errors.New("cannot peek from an empty queue")
	}

	return q.Data[0], nil

}

func (q *Queue) IsEmpty() bool {
	return len(q.Data) == 0
}
