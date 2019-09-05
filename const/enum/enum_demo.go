package main

func enums() {
	const (
		cpp = 0
		java
		python
		golang
	)
}

//...
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

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
	println(b)
	println(kb)
}
