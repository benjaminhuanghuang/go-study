package gostudy

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Open = 1 << iota
	Close
	Pending
)

func TestConstWeek(t *testing.T) {
	t.Log(Monday, Tuesday)
}
