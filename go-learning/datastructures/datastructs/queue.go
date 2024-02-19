package datastructs

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	removedItem := q.items[0]
	q.items = q.items[1:]
	return removedItem
}

func (q Queue) String() string {
	finalString := ""
	for _, item := range q.items {
		finalString += fmt.Sprintf("%d ", item)
	}
	return finalString
}
