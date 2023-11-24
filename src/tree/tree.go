package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	directoryCount = 0
	fileCount      = 0
)

func visit(path string, info os.FileInfo, err error) error {

	// When permissions deny tree walking inside a folder
	// it'll continue to the next folder.
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if info.IsDir() {
		fmt.Printf("%s|-- %s\n", getIndent(), info.Name())
		directoryCount++
	} else {
		fmt.Printf("|	%s|-- %s\n", getIndent(), info.Name())
		fileCount++
	}

	return nil
}

func getIndent() string {
	return ""
}

func main() {
	root := os.Args

	if len(os.Args) < 2 {
		root = append(root, "./")
	}
	err := filepath.Walk(root[1], visit)

	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", root, err)
	}

	fmt.Printf("\n%d directories, %d files\n", directoryCount, fileCount)
}
