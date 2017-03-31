package main

import (
	"os"
	"fmt"
	"strings"
	"path/filepath"
)

func main() {
	const (
		empty = ""
		separator = "---------------------------"
	)
	fmt.Println("Environment variables")
	fmt.Println(separator)
	fmt.Println(strings.Join(os.Environ(), "\n"))
	fmt.Println(empty)
	fmt.Println("Filesystem")
	fmt.Println(separator)
	_ = filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
		fmt.Printf("%s isdir:%t mode:%s modtime:%s size:%d\n", path, info.IsDir(), info.Mode().String(), info.ModTime().String(), info.Size())
		return nil
	})
}
