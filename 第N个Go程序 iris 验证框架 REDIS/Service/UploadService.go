package Service

import (
	"github.com/kataras/iris"
	"os"
	"io"
	"../Entity"
	"../upload"
)

func Upload(ctx iris.Context) {
	Name := ctx.PostValue("Name")
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
type UploadPage struct {

}
func (s *UploadPage)Get(ctx iris.Context) {
	upload.Writer(ctx)
}