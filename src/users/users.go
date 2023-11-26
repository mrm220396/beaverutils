/*

Package beaverutils/users returns the current user's home diretory, UID,
name and username. This library compounds the core of user commands for
beaverutils, this current file countains a type called SystemUser the same
as user.User from package os/user of the Golang's standard library.
The advantage of creating a new type as an alias of the standard library is
the possibility of including exclusives methods that are not included in the
standard library.

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

package users

import (
	"os"
	"os/user"
	"strings"
)

type SystemUser user.User

func (systemUser SystemUser) GetHomeDir() string {
	return systemUser.HomeDir
}

func (systemUser SystemUser) GetUID() string {
	return systemUser.Uid
}

func (systemUser SystemUser) GetUserName() string {
	return systemUser.Username
}

func (systemUser SystemUser) GetName() string {
	return systemUser.Name
}

func (systemUser SystemUser) GetGid() string {
	return systemUser.Gid
}

func Run() {
	lineargs := os.Args

	for _, v := range lineargs {
		v = strings.ToLower(v)
		if v == "--help" {
			i := info{}
			i.GetDescription()
		}
	}
}
func ReturnUsers() (SystemUser, error) {

	usr, err := user.Current()

	if err != nil {
		return SystemUser{}, err
	}

	response := SystemUser{
		Name:     usr.Name,
		Username: usr.Username,
		HomeDir:  usr.HomeDir,
		Gid:      usr.Gid,
		Uid:      usr.Uid,
	}
	return response, nil
}
