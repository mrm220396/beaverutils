package main

import (
	"beaverutils/src/users"
	"fmt"
)

func main() {
	user, _ := users.ReturnUsers()

	fmt.Println(user.Name, user.Username)
}
