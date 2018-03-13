package Controller

import (
	"github.com/kataras/iris"
	"../Service"
)
var registerController = Service.Register{}
func Register(ctx iris.Context) {
	registerController.Get(ctx)
}
