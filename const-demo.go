package gostudy

import "fmt"
import "math"

const s string = "constant"

func const_demo() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n

	// Constant expressions perform arithmetic with
	// arbitrary precision.
	fmt.Println(d)

	// A numeric constant has no type until it's given
	// one, such as by an explicit cast.
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
