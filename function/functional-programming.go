package funcstudy

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// apply has 3 parameters an return int
// 1. func, func has 2 parameter and return 2 int values
// 2. a
// 3. b
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)", opName, a, b)
	return op(a, b)
}

func main() {
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))
}
