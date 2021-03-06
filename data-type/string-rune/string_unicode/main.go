package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Go语言"

	// len() reutrns length of bytes, NOT the count of chars
	fmt.Println(len(s))

	fmt.Printf("%s\n", []byte(s)) // print the sentence
	fmt.Printf("%X\n", []byte(s)) // print the bytes

	// Print each byte, UTF-8 3bytes／汉字
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	// char/rune count
	count := utf8.RuneCountInString(s)
	fmt.Println(count)

	// byte -> unicode
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// for loop by runes
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()

}
