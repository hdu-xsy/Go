package DAO

import(
	"../Entity"
)
type AdminUser struct {

}
func (d *AdminUser) Get(admin Entity.AdminUser) (bool,error,Entity.AdminUser){
	bo, err := orm.Get(&admin)
	return bo,err,admin
}