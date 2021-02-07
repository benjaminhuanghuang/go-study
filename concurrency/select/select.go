package main

import (
	"fmt"
	"strconv"
	"time"
)

func sample1(ch chan string) {
	for i := 0; i < 19; i++ {
		ch <- "From sampl1" + strconv.Itoa(i)
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func sample2(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func main() {
	ch1 := make(chan string, 3)
	ch2 := make(chan int, 3)

	for i := 0; i < 10; i++ {
		go sample1(ch1)
		go sample2(ch2)
	}

	// Select the channel has data
	for {
		select {
		case str, ch1Check := <-ch1:
			if !ch1Check {
				fmt.Println("ch1 failed")
			}
			fmt.Println(str)
		case str, ch2Check := <-ch1:
			if !ch2Check {
				fmt.Println("ch1 failed")
			}
			fmt.Println(str)
		}
	}

	time.Sleep(10 * time.Second)
}
