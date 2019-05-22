package funcstudy

import (
	"fmt"
	"testing"
	"time"
)

//
func timeSpan(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent", time.Since(start).Seconds())
		return ret
	}
}

// defer

func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println("Clear resource")
	}()

	fmt.Println("Start")

	panic("err") // throw error

	fmt.Println("End")
}
