package queue

// use slice inernal
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v) // change the conent of the q
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
	// return head.(int)   // cast valuet to int
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
