package gostudy

import "fmt"

const (
	C0 = iota
	C1 = iota
	C2 = iota
)

//This can be simplified to
const (
	C = iota
	CC
	CCC
)
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

func test() {
	fmt.Println(C0, C1, C2) // "0 1 2"

}
