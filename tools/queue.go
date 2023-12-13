package queue

type Queue struct {
	Characters []rune
}

func (q *Queue) Len() int {
	return len(q.Characters)
}

func (q *Queue) Push(x any) {
	q.Characters = append(q.Characters, x.(rune))
}

func (q *Queue) Pop() any {
	old := q.Characters
	n := len(old)
	item := old[n-1]
	q.Characters = old[0 : n-1]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.Characters) == 0
}

func (q *Queue) Size() int {
	return len(q.Characters)
}

func NewQueue() *Queue {
	return &Queue{}
}
