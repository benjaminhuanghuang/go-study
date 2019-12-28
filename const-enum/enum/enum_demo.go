/*
enum 是一个const block

enum 可以用iota自增

可以用_skip某个值


Reference
	https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3
*/
package main

// Declare a new custom type: Weekday
type Weekday int

//...
const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// Timezone enum
type Timezone int

const (
	// iota: 0, EST: -5
	EST Timezone = -(5 + iota)
	// iota: 1, CST: -6
	CST
	// iota: 2, MST: -7
	MST
	// iota: 3, MST: -8
	PST
)

/***********************************************
Bitwise operations
***********************************************/
type Month int

const (
	January  Month = 1 << iota // 1 << 0 ==> 1
	February                   // 1 << 1 ==> 2
	March                      // 1 << 2 ==> 4
	April                      // 1 << 3 ==> 8
	May                        // 1 << 4 ==> 16
	June                       // ...
	July
	August
	September
	October
	November
	December
	AllMonths = January | February |
		March | April | May | June |
		July | August | September |
		October | November |
		December
)

func enums() {
	const (
		cpp = 0
		_
		java
		python
		golang
	)
}

/*
	Using iota
*/
const (
	Open = 1 << iota
	Close
	Pending
)

/*
	Using iota
*/
const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
)

func main() {
	println(Open)
	println(b, kb, mb, gb, tb, pb)
	println(AllMonths)
}
