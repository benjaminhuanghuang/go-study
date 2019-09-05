/*
定义变量有3种方法
var a int
var a
a := 10
*/
package main

import "fmt"

func variableZeroValue() {
	// variables have init value by default
	// This behavior is defferent with c/c++
	var a int
	var s string

	//fmt.Println(a, s)    //can not show empty string
	fmt.Printf("%d, %q\n", a, s) // %q means quotation, adding "" to string
}

func variableInitialValue() {
	var a, b int = 1, 2
	var s string = "Hello"

	fmt.Println(a, b, s)
}

func variableInitialValue2() {
	// auto inferred
	var a = 1
	var s = "Hello"

	fmt.Println(a, s)
}

/*
不能在function以外使用 :=
syntax error: non-declaration statement outside function bodygo
*/
//errorVariable := 1234

func variableShorter() {
	// auto
	a, b, c, s := 3, 4, true, "def"

	fmt.Println(a, b, c, s)
}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableShorter()
}
