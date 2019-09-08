package main

import (
	"bufio"
	"fmt"
	"os"

	"./fib"
)

func tryDefer() {
	// defer as a stack
	defer fmt.Println(1) // print after tryDefer return
	defer fmt.Println(2)
	fmt.Println(3)
	panic("Error occurred")
	fmt.Println(4)
}

func writeFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// file io buffer
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}
