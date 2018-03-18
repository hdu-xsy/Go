package Controller

import (
	"github.com/kataras/iris"
	"../Service"
)
var registerService = Service.Register{}
func Register(ctx iris.Context) {
	registerService.Get(ctx)
}

var deleteService = Service.Delete{}
func Delete(ctx iris.Context) {
	deleteService.Get(ctx)
}

var modifyService = Service.Modity{}
func Modify(ctx iris.Context) {
	modifyService.Get(ctx)
}