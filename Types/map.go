package main

import "fmt"

func main() {
	// create map <stirng:int>
	m := make(map[string]int)

	m["fool"] = 1234

	// iteration
	for key, value := range m {
		fmt.Printf("Key %s with value %d\n", key, value)
	}

	value, ok := m["abc"]
	if ok {
		fmt.Println("Got it", value)
	}
}
