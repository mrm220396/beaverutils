package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
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

	switch strings.ToLower(os.Args[1]) {
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
		wrong := fmt.Sprintf("Unfortunatly this argument is invalid\n%v\n\nWritten by: %v\n%v", Help, writtenby, GNU)
		return wrong
	}
}
