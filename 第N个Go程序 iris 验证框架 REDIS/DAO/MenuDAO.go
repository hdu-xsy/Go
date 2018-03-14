package DAO

import "../Entity"
type MenuDAO struct {

}

func (d *MenuDAO)Get(menu Entity.Menu) (bool,error,Entity.Menu) {
	bo,err:=orm.Get(&menu)
	return bo,err,menu
}
