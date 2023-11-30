/* Go tree - Print Directory trees


Copyright (C) 2023 Mateus R. Moreira
GitHub: https://www.github.com/mrm220396

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

Equivalent to 'id -un' in the GNU coreutils.
*/

package tree

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	directoryCount = 0
	fileCount      = 0
	root           []string
)

// GetFIleDoubleQuote responds to  -Q on command line, and returns
// Quote the names of files in double quotes.
func GetFileDoubleQuote(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("No such a file or directory %d", err)
		return nil
	}

	if info.IsDir() {
		fmt.Printf("%s|-- \"%s\"\n", getIndent(), info.Name())
	} else {
		fmt.Printf("	%s|-- \"%s\"\n", getIndent(), info.Name())
	}

	return err
}

// GetDirOnly returns only directory names
func GetDirOnly(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if info.IsDir() {
		fmt.Printf("%s|-- %s\n", getIndent(), info.Name())
		directoryCount++
	}

	return err
}

// Shows file and directories
func FileAndDir(path string, info os.FileInfo, err error) error {

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

func checkCommand(root []string) {
	switch root[1] {
	case "-d":
		filepath.Walk(root[2], GetDirOnly)
		break
	case "-Q":
		filepath.Walk(root[2], GetFileDoubleQuote)
		break
	}
}

// DefaultPrint returns the same as tree from CoreUtils
func DefaultPrint() {
	root = os.Args

	if len(os.Args) < 2 {
		root = append(root, "./")
	} else if strings.Contains(os.Args[1], "-") {
		root = append(root, "./")
		checkCommand(root)
	}
	err := filepath.Walk(root[1], FileAndDir)

	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", root, err)
	}

	fmt.Printf("\n%d directories, %d files\n", directoryCount, fileCount)
}
