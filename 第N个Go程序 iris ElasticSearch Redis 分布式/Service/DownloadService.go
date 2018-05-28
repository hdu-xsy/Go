package Service

import (
	"github.com/kataras/iris"
	"../Entity"
	"strconv"
)
type DownloadPage struct {

}
func (c *DownloadPage) Get(ctx iris.Context) {

	var FileList []Entity.File
	FileList = filedao.GetAll()
	ctx.ViewData("fileList",FileList)
	ctx.View("download.html")
}
func DownloadFile(ctx iris.Context) {
	name := ctx.Params().Get("Name")
	file := "./Files/"+name
	id,_ := strconv.ParseInt(name,10,64)
	ctx.SendFile(file,"请勿用于商业途径_" +filedao.GetName(id))
}