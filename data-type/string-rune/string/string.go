package stringstudy

import "fmt"

func main() {
	// Create
	book := "The color of magic"
	fmt.Println(book)

	// String length
	fmt.Println(len(book))

	// Access char in string
	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0]) //book[0] = 84 (type uint8)
	// uint8 is a byte

	// Strings in Go are immutable
	//book[0] = 123 // ERROR!

	// Slice string
	s := book[4:11]
	s = book[4:]
	s = book[:11]

	// concatenate strins
	s = "t" + book[1:]

	fmt.Println(s)

	// multiline
	poem := `Hello
	world
	...
	`
	fmt.Println(poem)
}
