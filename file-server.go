package gostudy

import (
	"net/http"
	"os"
)

func fileserver_demo() {
	// return dir and error
	dir, _ := os.Getwd()
	http.ListenAndServe(":3000", http.FileServer(http.Dir(dir)))
}
