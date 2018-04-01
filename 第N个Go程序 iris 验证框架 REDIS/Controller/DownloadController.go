package Controller

import (
	"github.com/kataras/iris"
	"../Download"
	"../DAO"
	"../Entity"
)
type DownloadController struct {

}

func (c *DownloadController) Get(ctx iris.Context) {
	var Filedao DAO.FileDAOInterface = new(DAO.FileDAO)
	var FileList []Entity.File
	FileList = Filedao.GetAll()
	Download.Writer(FileList,ctx)
}