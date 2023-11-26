package users

import "fmt"

var help = `


***Program inspired in users from coreutils. 
Originally written by Richard Mlynarik. Original C version of coreutils
***
Usage: users
Print the user name associated with the current effective user ID.
Same as id -un.

	--help        display this help and exit
	--version     output version information and exit
	users -a [OPTION] Shows the option to all users
	
	users home	   Shows the home directory
	users gid	   Shows the users' gid
	users name	   Shows the name
	users username Shows the username
	users anames   Shows a list of all usernames and names.
`

func Help() {
	fmt.Println(help)
}
