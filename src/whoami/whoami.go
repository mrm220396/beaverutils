/*
	LICENSE GPL

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
package whoami

import (
	"fmt"
	"os"
	"os/user"
)

var Version string = `0.0.01`

var writtenby string = "Mateus R. Moreira\nhttps://www.github.com/mrm220396"

var Help string = `
***Program inspired in whoami from coreutils. 
Originally written by Richard Mlynarik. Original C version of coreutils
***
Usage: whoami [OPTION]...
Print the user name associated with the current effective user ID.
Same as id -un.

	--help        display this help and exit
	--version     output version information and exit
	home 			output the home directory of your user
	username		output the current username
	  
	uid			output Uid is the user ID. On POSIX systems, 
	  			this is a decimal number representing the uid. 
	  			On Plan 9, this is the contents of /dev/user.
	  
	gid			Gid is the primary group ID. On POSIX systems, 
	  			this is a decimal number representing the gid.       			
`

var GNU string = `
Original GNU source:
GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Report any translation bugs to <https://translationproject.org/team/>
Full documentation <https://www.gnu.org/software/coreutils/whoami>
or available locally via: info '(coreutils) whoami invocation'

Written by Richard Mlynarik
`

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
