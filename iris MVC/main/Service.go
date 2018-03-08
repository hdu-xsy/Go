package main

import (
	"../Util"
	"../userlist"
)
func FindAllUser() []userlist.UserData{
	orm := Util.GetOnlineUser(*app)
	var userList []userlist.UserData
	orm.Find(&userList)
	return userList
}