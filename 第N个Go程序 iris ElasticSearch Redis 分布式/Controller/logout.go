package Controller

import (
	"../Service"
	"github.com/kataras/iris"
)

var logout = Service.Logout{}

type AdminLogout struct {

}
type Logout struct {

}
func (c *AdminLogout) Get(ctx iris.Context){
	logout.AdminLogout(ctx)
}
func (c *Logout) Get(ctx iris.Context){
	logout.Logout(ctx)
	uid = -1
}