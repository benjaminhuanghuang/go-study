package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 //
	fmt.Println(freq)
}
