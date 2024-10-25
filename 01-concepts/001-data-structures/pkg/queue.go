package pkg

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
	fmt.Println("Element added:", item)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return 0
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
