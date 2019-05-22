package main

import "fmt"

func main() {
	stocks := map[string]float64{
		"Amazon": 1699.8,
		"Google": 1699.8,
		"MSFT":   1699.8, // Must have trailing comma
	}

	value, ok := stocks["TSLA"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("TSLA not found")
	}

	// Delete
	delete(stocks, "Amazon")

	// iteration 1
	for key := range stocks {
		fmt.Println(key)
	}

	// create map <stirng:int>
	m := make(map[string]int)

	m["fool"] = 1234

	// iteration
	for key, value := range m {
		fmt.Printf("Key %s with value %d\n", key, value)
	}

	value2, ok2 := m["abc"]
	if ok2 {
		fmt.Println("Got it", value2)
	}
}
