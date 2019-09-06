package main

import "fmt"

func createMap() {
	m1 := map[string]float64{
		"Amazon": 1699.8,
		"Google": 1699.8,
		"MSFT":   1699.8, // Must have trailing comma
	}

	m2 := make(map[string]int) //m2 == empty map

	var m3 map[string]int // m3== nil but can be used as empty map

	fmt.Println(m1, m2, m3)
}

func traverseMap() {
	m := map[string]float64{
		"Amazon": 1699.8,
		"Google": 1699.8,
		"MSFT":   1699.8, // Must have trailing comma
	}

	// traversing
	for key, value := range m {
		fmt.Printf("Key %s with value %f\n", key, value)
	}
}

func main() {
	stocks := map[string]float64{
		"Amazon": 1699.8,
		"Google": 1699.8,
		"MSFT":   1699.8, // Must have trailing comma
	}

	// Getting values
	// Key 不存在时，获得value 类型的初始值
	if value, ok := stocks["TSLA"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("TSLA not found")
	}

	// Getting values 2
	value2, ok2 := stocks["TSLA"]
	if ok2 {
		fmt.Println("Got it", value2)
	}

	// Delete
	delete(stocks, "Amazon")

	// iteration 1
	// Key 无序
	for key := range stocks {
		fmt.Println(key)
	}

	// iteration 2
	for key, value := range stocks {
		fmt.Printf("Key %s with value %f\n", key, value)
	}
}
