package stringstudy

import "fmt"

func test() {
	sl := []string{"hello"}

	// append element
	sl = append(sl, " go")

	fmt.Println(sl)

	fmt.Println(sl[0:1])

}
