/*
A Tour of Go Jun 20, 2012
https://www.youtube.com/watch?time_continue=19&v=ytEkHepK08c
*/
package main

import "fmt"

type Office int

const (
	Boston Office = iota
	NewYork
)

var officePlace = [...]string{
	"Boston",
	"NewYork",
}

// Implement interface
func (o Office) String() string {
	return "Google, " + officePlace[o]
}

func main() {
	fmt.Printf("Hello, %s\n", Boston)
}
