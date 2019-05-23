package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes春眠不觉晓"
	fmt.Println(len(s)) // len() reutrns length of bytes, NOT the count of chars
	//RuneCountInString(s) returns char count

	fmt.Printf("%s\n", []byte(s)) // print the sentence

	fmt.Printf("%X\n", []byte(s)) // print the bytes

	// Print bytes, 每个中文字符三字节
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b) // print the digits
	}
	fmt.Println()

	// Print runes
	for _, ch := range s {
		fmt.Printf("%X ", ch) // print the digits
	}
	fmt.Println()

	// Decode bytes to runes
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes) // 每次转码一个rune
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// Print runes with index
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
}
