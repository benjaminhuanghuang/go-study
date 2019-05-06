/*
- 必须是 main package
- 必须是 main() method
- 文件名不一定是 main.go

- main() method 没有返回值, using os.Exit(1) return status.

- main() 不支持传入参数，using os.Args
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	// get args count
	argCount := len(os.Args[1:])
	fmt.Printf("Totally %d argument", argCount)

	n := 1
	s := fmt.Sprintf("%d", n)
	fmt.Printf("s = %s (type %T)\n", s, s)

	os.Exit(0)
}
