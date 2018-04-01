package Controller

import (
	"github.com/kataras/iris"
	"../Entity"
	"../upload"
	"os"
	"io"
	"../DAO"
)
type UploadController struct{
}
func (c *UploadController) BeginRequest(ctx iris.Context) {
	if auth, _ := Entity.Sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.Redirect("/adminlogin")
		return
	}
}
func (c *UploadController) EndRequest(ctx iris.Context) {}
func (c *UploadController) Get(ctx iris.Context) {
	upload.Writer(ctx)
}

func Uploads(ctx iris.Context) {
	Name := ctx.PostValue("Name")
	var filedao DAO.FileDAOInterface = new(DAO.FileDAO)
	filedao.Insert(Entity.File{Name:Name})
	file, info, err := ctx.FormFile("uploadfile")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	defer file.Close()
	fname := info.Filename
	out, err := os.OpenFile("./Files/"+fname,
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	defer out.Close()
	io.Copy(out, file)
	ctx.Redirect("/backend")
}