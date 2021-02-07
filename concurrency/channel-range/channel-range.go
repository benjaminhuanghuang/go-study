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
	// if we do not close the channel, read empty channel will cause an error
	close(channel)
}

func main() {
	/*
	 create channel can have 3 string
	*/
	var messageChannel = make(chan string, 3)

	go sample1(messageChannel) // put message 1, 2, 3 to the channel
	go sample2(messageChannel) // read message 1 -> message 4 enter channel -> routine2 message enter channel

	time.Sleep(3 * time.Second)

	for str := range messageChannel {
		// for loop breads when channel is closed
		// if we do not close the channel, read empty channel will cause an error
		fmt.Println(str)
	}

	fmt.Println("Hello from --- main()")
}
