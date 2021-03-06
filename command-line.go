/*
	go run command-line.go -port=5151 -path=site

	go run command-line.go -help

*/
package gostudy

import (
	"flag"
	"net/http"
	"os"
)

func main() {
	var dir string

	port := flag.String("port", "3000", "port to serve HTTP on")
	path := flag.String("path", "", "path to server")
	flag.Parse()

	if *path == "" {
		dir, _ = os.Getwd()
	} else {
		dir = *path
	}
	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))
}
