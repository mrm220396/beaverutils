/* 2023

Package users based on users from coreutils
Under GLP 3.0

*/

package users

import (
	"os/user"
)

type SystemUser struct {
	Name     string
	Username string
	HomeDir  string
	Gid      string
	UID      string
}

func (systemUser SystemUser) GetHomeDir() string {
	return systemUser.HomeDir
}

func (systemUser SystemUser) GetUID() string {
	return systemUser.UID
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
		UID:      usr.Uid,
	}
	return response, nil
}
