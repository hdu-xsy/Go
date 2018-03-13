package DAO

import (
	"../Entity"
)
type UserData struct {

}

func (d *UserData)FindAll() []Entity.UserData {
	var userList []Entity.UserData
	orm.Find(&userList)
	return userList
}
func (d *UserData) Get(userdata Entity.UserData) (bool,error,Entity.UserData){
	bo, err := orm.Get(&userdata)
	return bo,err,userdata
}
func (d *UserData) Insert(userdata Entity.UserData) (int64,error) {
	i,err := orm.Insert(&userdata)
	return i,err
}
