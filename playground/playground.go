package main

import "fmt"

// const s string = "constant"

func main() {
	// Create
	book := "The color of magic"
	fmt.Println(book)

	// String length
	fmt.Println(len(book))

	// Access char in string
	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
}
