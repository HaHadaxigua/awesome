package main

import (
	"os"
	"path/filepath"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	//f, err := os.OpenFile(filepath.Join(path, "a/b"), os.O_CREATE|os.O_RDWR, os.ModeDir)
	f, err := os.OpenFile(filepath.Join(path, "a/b/c.txt"), os.O_CREATE|os.O_RDWR, os.ModeDir)
	if err != nil {
		panic(err)
	}
	f.Close()
}
