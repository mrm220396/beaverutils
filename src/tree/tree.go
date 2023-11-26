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
)

func GetDirOnly() {

}

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

func checkCommand(root []string) {
	fmt.Println("Sometime soon")
}

func DefaultPrint() {
	root := os.Args

	if len(os.Args) < 2 {
		root = append(root, "./")
	} else if strings.Contains(os.Args[1], "--") {
		checkCommand(root)
	}
	err := filepath.Walk(root[1], visit)

	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", root, err)
	}

	fmt.Printf("\n%d directories, %d files\n", directoryCount, fileCount)
}
