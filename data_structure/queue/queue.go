package queue

// use slice inernal
type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v) // change the conent of the q
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
