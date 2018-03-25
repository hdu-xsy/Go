package Controller

import (
	"github.com/kataras/iris"
	"../Service"
)
var CommentService = Service.Comment{}
func CommentController(ctx iris.Context) {
	CommentService.Get(ctx)
}
