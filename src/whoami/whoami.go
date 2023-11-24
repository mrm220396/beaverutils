/* Go Whoami whoami - Print information about the current user in Go.


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

package main

import (
	"fmt"
	"os"
	"os/user"
)

var writtenby string = "Mateus R. Moreira\nhttps://www.github.com/mrm220396"

func main() {
	resp := WhoamI()
	fmt.Println(resp)
}

func WhoamI() string {
	// Current user receives all categories listed to user.Current()
	// from the package os/user
	currentUser, err := user.Current()
	if err != nil {
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		return currentUser.Name
	}

	switch os.Args[1] {
	case "--version":
		return Version
	case "--help":
		return Help
	case "uid":
		return currentUser.Uid
	case "home":
		return currentUser.HomeDir
	case "gid":
		return currentUser.Gid
	case "username":
		return currentUser.Username

	case "--original":
		return GNU

	default:
		wrong := fmt.Sprintf("Unfortunatly this argument is invalid\n%v\n\nWritten by: %v\n%v",
			Help, writtenby, GNU)
		return wrong
	}
}
