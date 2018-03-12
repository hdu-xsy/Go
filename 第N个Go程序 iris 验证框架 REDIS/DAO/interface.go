package DAO

import(
	"github.com/kataras/iris"
	"../Entity"
	"../Util"
)
var app = iris.New()
var orm = Util.Getorm(*app)
var err1 = Util.GetUserData(*app)
var err2 = Util.GetAdminUser(*app)
var err3 = Util.GetOnlineUser(*app)

type UserDataDAOInterface interface {
	FindAll() []Entity.UserData
	Get(userdata Entity.UserData) (bool,error,Entity.UserData)
}

type AdminUserDAOInterface interface {
	Get(admin Entity.AdminUser) (bool,error,Entity.AdminUser)
}

type OnlineUserDAOInterface interface {
	Get(onlineuser Entity.OnlineUser) (bool,error,Entity.OnlineUser)
	Insert(onlineuser Entity.OnlineUser) error
	Delete(onlineuser Entity.OnlineUser) error
}
