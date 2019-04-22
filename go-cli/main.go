package main

import (
	"fmt"
	"os"
)

func main() {
	// get args count
	argCount := len(os.Args[1:])
	fmt.Printf("Totally %d argument", argCount)
}
