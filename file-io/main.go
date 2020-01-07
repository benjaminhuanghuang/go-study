package main

import (
	"io/ioutil"
	"os"
	"path"
)

func ReadFileInUserHome() ([]byte, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	filePath := path.Join(home, "Desktop", "leetcode.html")

	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func main() {
	return
}
