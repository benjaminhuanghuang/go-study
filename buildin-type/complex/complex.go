package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func modulus() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c)) // Complex modulus

}

func euler() {
	//fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Printf("%.3f", cmplx.Exp(1i*math.Pi)+1)
}

func main() {
	euler()
}
