package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// file io buffer
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, "aaaaaaaaaaaaa")
	}
}

func main() {
	writeFile("test.txt")
}
