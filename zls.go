package main

import (
	"fmt"
	"os"
	"path/filepath"
	"log"
)

type FileTree struct {
	path     string
	info     *os.FileInfo
	children []*FileTree
}

func listFiles(path string) ([]FileTree, error) {
	var files []FileTree
	var parseEntries = func(path string, info os.FileInfo, err error) error {
		files = append(files, FileTree{
			path: path,
			info: &info,
			children: []*FileTree{},
		})
		return nil
	}
	err := filepath.Walk(path, parseEntries)
	return files, err
}

func printTree(tree []FileTree) {
	reset := string("\033[00m")
	blue  := string("\033[34m")

	for _, f := range tree {
		fmt.Println(f.info.Name)
		fmt.Println(blue, f.info.Name, reset)
	}
}

func main() {
	var files []FileTree
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	files, err = listFiles(path)
	if err != nil {
		log.Println(err)
	}
	printTree(files)
}
