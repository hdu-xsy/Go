package DAO

import "../Entity"

type OnlineUser struct {

}
func (d *OnlineUser) Get(onlineuser Entity.OnlineUser) (bool,error,Entity.OnlineUser){
	bo, err := orm.Get(&onlineuser)
	return bo,err,onlineuser
}
func (d *OnlineUser) Insert(onlineuser Entity.OnlineUser) error{
	_,err:= orm.Insert(&onlineuser)
	return err
}
func (d *OnlineUser) Delete(onlineuser Entity.OnlineUser) error{
	_,err := orm.Delete(&onlineuser)
	return err
}