/*

GO pwd - Prints the current folder the user is in

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

package pwd

import (
	"fmt"
	"os"
)

func Pwd() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(currentDir)
}
