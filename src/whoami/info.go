package main

var Version string = `0.0.01`


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