package main

import (
	"fmt"
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
	// dir, err := os.Executable()
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	dir := os.Args[0]
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(dir)

	fmt.Println(">> Current Path: " + dir)
}
