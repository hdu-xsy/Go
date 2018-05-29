package DAO

import "../Entity"
type MenuDAO struct {

}

func (d *MenuDAO)Get(menu Entity.Menu) (bool,error,Entity.Menu) {
	bo,err:=orm.Get(&menu)
	return bo,err,menu
}
func (d *MenuDAO)GetAll() []Entity.Menu{
	var menu []Entity.Menu
	orm.Asc("id").Find(&menu)
	return menu
}
func (d *MenuDAO) Insert(menu Entity.Menu) (int64,error) {
	i,err := orm.Insert(&menu)
	return i,err
}
func (d *MenuDAO) Delete(menu Entity.Menu) (int64,error) {
	i,err := orm.Id(menu.Id).Delete(&menu)
	return i,err
}
func (d *MenuDAO) Motify(menu Entity.Menu) (int64,error) {
	i,err := orm.Id(menu.Id).Update(&menu)
	return i,err
}