/*

 */

package main

import (
	"fmt"
	"time"
)

func sample1(channel chan string) {
	channel <- "hello from go routine 1"
	channel <- "hello from go routine 2"
	channel <- "hello from go routine 3"
	channel <- "hello from go routine 4"
}

func sample2(channel chan string) {
	time.Sleep(2 * time.Second) // make sample2 runs after sample 1
	str := <-channel
	str = str + "This is go routine 2"
	channel <- str
}

func main() {
	/*
		create channel can have 3 string
	*/
	var messageChannel = make(chan string, 3)

	go sample1(messageChannel) // put message 1, 2, 3 to the channel
	go sample2(messageChannel) // read message 1 -> message 4 enter channel -> routine2 message enter channel

	time.Sleep(3 * time.Second)

	fmt.Println(<-messageChannel) // message 2
	fmt.Println(<-messageChannel) // message 3
	fmt.Println(<-messageChannel) // message 4
	fmt.Println(<-messageChannel) // routine2 message

	// cause dead lock because all goroutines are asleep
	// fmt.Println(<-messageChannel)

	fmt.Println("Hello from --- main()")
}
