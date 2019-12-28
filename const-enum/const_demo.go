package main

import (
	"fmt"
	"math"
)

const (
	a, b = 3, 4
)

func calculateWithoutCast() {
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func calculate() {
	a := 3
	b := 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	calculate()
}
