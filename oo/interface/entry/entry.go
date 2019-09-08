package main

import (
	"fmt"
	"time"

	"../interfaces"
	"../mock"
	"../real"
)

func download(r interfaces.Retriever) string {
	return r.Get("http://www.google.com")
}

func post(poster interfaces.Poster) {
	poster.Post("http://www.google.com", map[string]string{
		"name": "ben",
	})
}

//
func session(s interfaces.RetrieverPoster) {
	s.Get("http://www.google.com")
	s.Post("http://www.google.com", map[string]string{
		"name": "ben",
	})
}

func inspect(r interfaces.Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
}

func main() {
	var r interfaces.Retriever
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
