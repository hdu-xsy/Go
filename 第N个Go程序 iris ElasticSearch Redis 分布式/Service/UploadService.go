package Service

import (
	"github.com/kataras/iris"
	"../Entity"
	"mime/multipart"
	"strconv"
)

func Upload(ctx iris.Context) {
	Name := ctx.PostValue("Name")
	filedao.Insert(Entity.File{Name:Name})
	ctx.UploadFormFiles("./Files",beforeSave)
	ctx.Redirect("/backend/1")
}
type UploadPage struct {

}
func (s *UploadPage)Get(ctx iris.Context) {
	ctx.View("upload.html")
}
func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	var filelist []Entity.File
	filelist = filedao.GetAll()
	file.Filename = strconv.Itoa(len(filelist))
}