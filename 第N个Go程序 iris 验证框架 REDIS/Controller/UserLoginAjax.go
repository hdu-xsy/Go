package Controller

import (
	"../Service"
	"github.com/kataras/iris"
)
var userloginajax = Service.UserLoginAjax{}
func UserLoginAjax(ctx iris.Context) {
	uid = userloginajax.Get(ctx)
}
