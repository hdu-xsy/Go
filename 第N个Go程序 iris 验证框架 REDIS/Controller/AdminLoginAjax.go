package Controller

import (
	"github.com/kataras/iris"
	"../Service"
)
var adminloginajax = Service.AdminLoginAjax{}
func AdminLoginAjax(ctx iris.Context) {
	adminloginajax.Get(ctx)
}
