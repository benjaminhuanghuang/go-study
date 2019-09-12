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

//
func tryDefer2() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close() // finally

	// file io buffer
	writer := bufio.NewWriter(file)
	defer writer.Flush() // finally

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	// writeFile("fib.txt")
	tryDefer2()
}
