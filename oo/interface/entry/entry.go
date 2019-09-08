package main

import (
	"fmt"
	"time"

	"../mock"
	"../real"
	"../retriever"
)

func download(r retriever.Retriever) string {
	return r.Get("http://www.google.com")
}

func inspect(r retriever.Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
}

func main() {
	var r retriever.Retriever
	r = mock.Retriever{Contents: "Thie is mock retriver"}
	inspect(r)
	// Print type and value
	fmt.Printf("%T %v\n", r, r)
	fmt.Println(download(r))

	/*********************************************
	// Type assertion
	*********************************************/
	if realRetriver, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriver.TimeOut)
	} else {
		fmt.Println("not a real retriver")
	}

	r = &real.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	fmt.Printf("%T %v\n", r, r)
	fmt.Println(download(r))
}
