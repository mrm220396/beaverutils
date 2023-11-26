// Copyright 2023 Beaverutils author Mateus Moreria, all rights reserved.
// Use of this source code is governed by a GPL 3.0 based license

package users

type info struct {
	//Attribuite name returns this package name
	Name string `json:"name"`
	// Version of the program stored in beaverutils/src/users/config.json under
	// the entry "version".
	Version string `json:"version"`
	// description receives a string of the instructions of the program
	// the attribute Description is used by beaverutils/users when receives
	// a wrong argument from the command line or the UNIX system use types
	// users --help.
	description string
	// Author returns author's name of this library
	// The original authors are listed at GNU coreutils.
	Author string `json:"author"`
}

func (i *info) GetDescription() string {
	i.description = `Usage: users
	Print the user name associated with the current effective user ID.
	Same as id -un.
	
		--help        display this help and exit
		--version     output version information and exit
		users -a [OPTION] Shows the option to all users
		
		users home	   Shows the home directory
		users gid	   Shows the users' gid
		users name	   Shows the name
		users username Shows the username
		users anames   Shows a list of all usernames and names.`
	return i.description
}
