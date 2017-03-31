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
	var skipDirs []string
	if env := os.Getenv("PLUGIN_SKIP_DIRECTORIES"); len(env) > 0 {
		skipDirs = strings.Split(env, ",")
	} else {
		skipDirs = make([]string, 0)
	}

	fmt.Println("Environment variables")
	fmt.Println(separator)
	fmt.Println(strings.Join(os.Environ(), "\n"))
	fmt.Println(empty)
	fmt.Println(empty)
	fmt.Println(empty)
	fmt.Println("Filesystem")
	fmt.Println(separator)
	_ = filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info == nil {
			return nil
		}

		if info.IsDir() {
			for _, skip := range skipDirs {
				if path == skip {
					return filepath.SkipDir
				}
			}
			fmt.Printf("%s isdir:%t mode:%q modtime:%q size:%d\n", path, info.IsDir(), info.Mode().String(), info.ModTime().String(), info.Size())
		}
		return nil
	})
}
