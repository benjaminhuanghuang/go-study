package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`

	words := strings.Fields(text)
	counts := map[string]int{} // Empty map

	// loop through
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	for key, value := range counts {
		fmt.Printf("Word %q with count %d\n", key, value)
	}
}
