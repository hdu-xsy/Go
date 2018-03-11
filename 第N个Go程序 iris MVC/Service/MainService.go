package Service

import (
	"../DAO"
	"github.com/kataras/iris"
)
var app = iris.New()
func checkError(err error) {
	if err != nil {
		app.Logger().Fatalf("err:",err)
	}
}
var uid int64
var userdatadao DAO.UserDataDAOInterface = new(DAO.UserData)
var adminuserdao DAO.AdminUserDAOInterface = new(DAO.AdminUser)
var onlineuserdao DAO.OnlineUserDAOInterface = new(DAO.OnlineUser)